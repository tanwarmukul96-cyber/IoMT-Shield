package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/tarwanmukul96-cyber/IoMT-Shield/gateway/internal/policy"
)

type MedicalTelemetry struct {
	Device    string `json:"device"`
	HeartRate int    `json:"heart_rate"`
	SpO2      int    `json:"spo2"`
}

const medicalServiceURL = "http://medical-service:9000"

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	_, _ = w.Write([]byte(`{"status":"healthy","service":"IoMT-Shield Gateway"}`))
}

func medicalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read incoming medical telemetry.
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request", http.StatusBadRequest)
		return
	}

	var telemetry MedicalTelemetry

	if err := json.Unmarshal(body, &telemetry); err != nil {
		log.Printf("[SECURITY] Invalid JSON telemetry: %v", err)
		http.Error(w, "Invalid telemetry format", http.StatusBadRequest)
		return
	}

	// Evaluate telemetry through the Policy Engine.
	decision := policy.Evaluate(policy.Telemetry{
		Device:    telemetry.Device,
		HeartRate: telemetry.HeartRate,
		SpO2:      telemetry.SpO2,
	})

	log.Printf(
		"[POLICY] Device=%s Action=%s Reason=%s",
		telemetry.Device,
		decision.Action,
		decision.Reason,
	)

	// BLOCK is the only decision that stops traffic here.
	if decision.Action == policy.Block {
		log.Printf(
			"[SECURITY] BLOCKED device=%s reason=%s",
			telemetry.Device,
			decision.Reason,
		)

		http.Error(
			w,
			"Blocked by IoMT-Shield policy",
			http.StatusForbidden,
		)
		return
	}

	// ALERT is deliberately fail-open:
	// medical traffic continues while the event is logged.
	if decision.Action == policy.Alert {
		log.Printf(
			"[SECURITY] ALERT device=%s reason=%s",
			telemetry.Device,
			decision.Reason,
		)
	}

	// Restore request body before forwarding.
	r.Body = io.NopCloser(
		bytes.NewReader(body),
	)

	// Forward approved telemetry to the medical service.
	target, err := url.Parse(medicalServiceURL)
	if err != nil {
		http.Error(w, "Medical service configuration error", http.StatusInternalServerError)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ErrorHandler = func(
		w http.ResponseWriter,
		r *http.Request,
		err error,
	) {
		log.Printf("[GATEWAY] Medical service unavailable: %v", err)
		http.Error(
			w,
			"Medical service unavailable",
			http.StatusBadGateway,
		)
	}

	proxy.ServeHTTP(w, r)
}

func main() {

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/medical", medicalHandler)

	addr := ":8080"

	log.Printf("IoMT-Shield Gateway listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
