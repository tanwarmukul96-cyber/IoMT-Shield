import time
import random
import requests

GATEWAY_URL = "http://gateway:8080/medical"

print("IoMT Patient Monitor Simulator started")

while True:
    heart_rate = random.randint(65, 95)
    oxygen = random.randint(95, 100)

    payload = {
        "device": "patient-monitor-01",
        "heart_rate": heart_rate,
        "spo2": oxygen
    }

    try:
        response = requests.post(
            GATEWAY_URL,
            json=payload,
            timeout=2
        )

        print(
            f"[PATIENT MONITOR] "
            f"HR={heart_rate} BPM | "
            f"SpO2={oxygen}% | "
            f"Gateway={response.status_code}"
        )

    except Exception as e:
        print(f"[GATEWAY ERROR] {e}")

    time.sleep(5)