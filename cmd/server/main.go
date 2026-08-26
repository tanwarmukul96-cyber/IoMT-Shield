package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "IoMT-Shield Gateway: HEALTHY")
}

func main() {
	http.HandleFunc("/health", healthHandler)

	addr := ":8080"
	log.Printf("IoMT-Shield Gateway listening on %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}