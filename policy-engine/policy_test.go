package policy

import "testing"

func TestEvaluate(t *testing.T) {

	tests := []struct {
		name     string
		input    Telemetry
		expected Action
	}{
		{
			name: "Valid telemetry",
			input: Telemetry{
				Device:    "patient-monitor-01",
				HeartRate: 80,
				SpO2:      98,
			},
			expected: Allow,
		},
		{
			name: "Missing device identity",
			input: Telemetry{
				Device:    "",
				HeartRate: 80,
				SpO2:      98,
			},
			expected: Block,
		},
		{
			name: "Invalid heart rate",
			input: Telemetry{
				Device:    "patient-monitor-01",
				HeartRate: 500,
				SpO2:      98,
			},
			expected: Alert,
		},
		{
			name: "Invalid oxygen level",
			input: Telemetry{
				Device:    "patient-monitor-01",
				HeartRate: 80,
				SpO2:      20,
			},
			expected: Alert,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			result := Evaluate(test.input)

			if result.Action != test.expected {
				t.Fatalf(
					"expected %s, got %s (%s)",
					test.expected,
					result.Action,
					result.Reason,
				)
			}
		})
	}
}
