# IoMT-Shield

IoMT-Shield is an open-source security gateway designed to protect legacy and connected **Internet of Medical Things (IoMT)** environments.

The project explores how vulnerable medical-device networks can be protected through an external security layer without requiring modifications to the medical device itself.

## Problem

Legacy medical devices may operate on unsupported operating systems or software stacks where security updates and direct modifications are difficult, risky, or operationally impractical.

This creates an exposed attack surface around healthcare infrastructure.

IoMT-Shield reduces this exposure by placing a security gateway between medical devices and backend services, allowing traffic and device telemetry to be monitored and evaluated before reaching protected systems.

## Solution

IoMT-Shield uses a gateway-based architecture that separates the medical-device environment from external networks.

The gateway currently provides:

* Telemetry validation
* Policy-based security decisions
* Abnormal-behaviour alerts
* Docker-based network isolation
* Structured security logging

The architecture is designed to evolve toward deeper network inspection, behavioural detection, and automated containment.

## Architecture

```text
External / Hospital Network
          |
          v
   IoMT-Shield Gateway
          |
          v
      Policy Engine
       /        \
   ALLOW       ALERT
     |           |
     v           |
Simulated        |
Medical Service  |
     ^           |
     |           |
Patient Monitor  |
   Simulator     |
```

## Current Capabilities

* External security gateway for IoMT environments
* Go-based gateway and policy engine
* Medical-device telemetry simulation
* JSON telemetry validation
* Device identity validation
* Heart-rate and SpO2 policy checks
* Security alerts for abnormal telemetry
* Docker-based network isolation
* Structured security logging
* Automated Go testing
* Reproducible local deployment

## Technology Stack

* Go
* Python
* Docker
* Docker Compose
* Linux / Alpine
* YAML
* REST APIs
* Git and GitHub

## Project Structure

```text
IoMT-Shield/
├── gateway/
├── devices/
├── policy-engine/
├── configs/
├── docker-compose.yml
├── LICENSE
└── README.md
```

## Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/tanwarmukul96-cyber/IoMT-Shield.git
cd IoMT-Shield
```

### 2. Start the environment

```bash
docker compose up --build
```

### 3. Check running services

```bash
docker compose ps
```

### 4. View gateway logs

```bash
docker compose logs gateway
```

### 5. Stop the environment

```bash
docker compose down
```

## Example Telemetry

A valid telemetry request:

```json
{
  "device": "patient-monitor-01",
  "heart_rate": 80,
  "spo2": 98
}
```

Normal telemetry is allowed to reach the simulated medical service.

Abnormal telemetry, such as an unrealistic heart-rate value, generates a security alert and is recorded by the gateway.

Example:

```text
[POLICY] Device=patient-monitor-01 Action=ALERT Reason=Heart-rate value outside configured range
```

## Testing

### Run the gateway tests

```bash
go test ./...
```

### Build the gateway

```bash
go build ./...
```

The current test environment uses simulated medical devices and services and is intended for controlled cybersecurity experimentation.

## Future Direction

Planned development areas include:

* DICOM and HL7-aware inspection
* TLS-aware traffic analysis
* Behavioural anomaly detection
* Machine-learning-based risk scoring
* Cross-node attack correlation
* MITRE ATT&CK mapping
* eBPF/XDP-based packet enforcement
* Dynamic network quarantine
* Micro-segmented device policies
* Performance and latency benchmarking
* Safety-aware automated response

These capabilities are part of the development roadmap and are **not represented as fully implemented in the current release**.

## Safety and Disclaimer

IoMT-Shield is a research and cybersecurity engineering project using simulated medical environments.

It is **not a certified medical device or clinical security product** and must not be deployed on real medical equipment or production healthcare infrastructure without appropriate safety, regulatory, and security validation.

## License

This project is licensed under the terms specified in the repository's `LICENSE` file.
