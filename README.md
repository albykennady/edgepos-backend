# EdgePOS Backend

EdgePOS is a learning project that simulates a modern retail
point-of-sale system using Go and PostgreSQL.

## Project Goals

- Build a REST API using Go.
- Persist products, carts and transactions in PostgreSQL.
- Implement store, terminal and cashier context.
- Model cart states and reliable checkout processing.
- Build an operator interface using React and TypeScript.
- Simulate cash and card payments.
- Study EMV payment architecture and terminal integration.
- Practice testing, Git workflows, CI/CD and containerization.

## Planned Technology Stack

- Backend: Go
- Database: PostgreSQL
- Frontend: React and TypeScript
- Local environment: Docker Compose
- Source control and CI: GitHub and GitHub Actions

## Current Status

Day 1: Repository and project foundation.

The Go API health endpoint is the next implementation milestone.

## Payment Safety

This project initially uses simulated payments for learning.
It does not process real card payments or store real payment card data.

A production EMV integration would require compatible certified
hardware, a supported payment provider and the applicable security
and certification requirements.


## Running Locally

### Prerequisites

- Go installed
- Git

### Start the API

From the project root, run:

```bash
go run ./cmd/api
```

The server listens on port 8080.

### Health Endpoint

Request:

```http
GET http://localhost:8080/health
```

Expected response:

- Status: `200 OK`
- Content-Type: `application/json`
- Body: `{"status":"ok"}`

### Run Tests

```bash
go test ./...
```

### Current Implementation

- [x] Go HTTP server
- [x] chi router
- [x] Health endpoint
- [x] Automated health endpoint test