package policy

type Action string

const (
	Allow      Action = "ALLOW"
	Alert      Action = "ALERT"
	Block      Action = "BLOCK"
	Restrict   Action = "RESTRICT"
	Quarantine Action = "QUARANTINE"
)

type Telemetry struct {
	Device    string
	HeartRate int
	SpO2      int
}

type Decision struct {
	Action Action
	Reason string
}

func Evaluate(t Telemetry) Decision {
	if t.Device == "" {
		return Decision{
			Action: Block,
			Reason: "Missing device identity",
		}
	}

	if t.HeartRate < 20 || t.HeartRate > 250 {
		return Decision{
			Action: Alert,
			Reason: "Heart-rate value outside configured range",
		}
	}

	if t.SpO2 < 50 || t.SpO2 > 100 {
		return Decision{
			Action: Alert,
			Reason: "SpO2 value outside configured range",
		}
	}

	return Decision{
		Action: Allow,
		Reason: "Telemetry passed safety policy",
	}
}
