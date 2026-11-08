# EdgePOS

A learning-focused Point of Sale (POS) backend project built with Go, with the goal of understanding how modern retail POS systems are designed and developed.

The project follows a practical software engineering workflow, covering API development, database integration, transaction management, payment simulation, testing, and containerization.

## Tech Stack

### Backend
- **Language:** Go
- **HTTP Router:** Chi
- **Logging:** Go `log/slog`
- **Database:** PostgreSQL (planned)
- **Database Driver:** pgx (planned)

### Frontend
- React (planned)
- TypeScript (planned)

### Development & Infrastructure
- Git and GitHub
- Docker and Docker Compose (planned)
- Unit and integration testing
- REST APIs

## Current Features

The following backend capabilities have been implemented:

- **Health Check:** Endpoint to verify that the API server is running.
- **Readiness Check:** Endpoint to indicate whether the application is ready to serve requests.
- **Environment Configuration:** Server port can be configured through environment variables.
- **Structured Logging:** JSON-based application logging using Go's `log/slog`.
- **Graceful Shutdown:** Handles termination signals and allows the HTTP server to shut down gracefully.
- **Automated Tests:** HTTP endpoint and configuration tests.

## API Endpoints

| Method | Endpoint | Description |
|---|---|---|
| GET | `/health` | Checks whether the API server is running. |
| GET | `/ready` | Checks whether the application is ready to serve requests. |

### Example Responses

**GET /health**

```json
{
  "status": "ok"
}
```

**GET /ready**

```json
{
  "status": "ready"
}
```

## Getting Started

### Prerequisites

- Go installed
- Git installed

### Clone the Repository

```bash
git clone https://github.com/albykennady/edgepos-backend.git
cd edgepos-backend
```

### Run the Application

```bash
go run ./cmd/api
```

By default, the API server starts on port `8080`.

### Configure the Server Port

You can override the default port using the `PORT` environment variable.

```bash
PORT=9090 go run ./cmd/api
```

The server will then listen on port `9090`.

## Running Tests

Run all tests:

```bash
go test ./...
```

Run static analysis:

```bash
go vet ./...
```

## Project Structure

```text
edgepos-backend/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── config/
│   │   └── config.go
│   │
│   └── httpapi/
│       ├── router.go
│       ├── router_test.go
│       ├── health.go
│       └── ready.go
│
├── go.mod
├── go.sum
└── README.md
```

### Directory Responsibilities

- `cmd/api`: Application entry point and server startup.
- `internal/config`: Application configuration and environment variable handling.
- `internal/httpapi`: HTTP router, endpoint handlers, and related tests.

## Development Roadmap

The project will progressively introduce POS functionality.

### Backend

- [x] HTTP server foundation
- [x] Health and readiness endpoints
- [x] Environment-based configuration
- [x] Structured logging
- [x] Graceful shutdown
- [ ] PostgreSQL integration
- [ ] Database migrations
- [ ] Store and terminal management
- [ ] Cashier sessions
- [ ] Product and SKU management
- [ ] Cart management
- [ ] Discounts and void operations
- [ ] Payment simulation
- [ ] Transaction management
- [ ] gRPC service design

### Frontend

- [ ] React and TypeScript application
- [ ] POS operator interface
- [ ] Product search and scanning simulation
- [ ] Cart and checkout interface
- [ ] Payment workflow

### Infrastructure

- [ ] Dockerize the backend
- [ ] Docker Compose development environment
- [ ] Integration testing
- [ ] CI/CD pipeline
- [ ] Cloud deployment exploration

## Project Goals

The main objective is to build a practical understanding of modern POS architecture and backend engineering.

The project will focus on:

- Writing maintainable and testable Go code.
- Designing REST APIs and understanding service architecture.
- Working with relational databases and transactions.
- Understanding concurrency and reliability in transaction-heavy systems.
- Learning containerization and deployment workflows.
- Exploring payment processing concepts and EMV architecture.

**Note:** This is an independent learning project. Payment functionality will initially be simulated; it is not connected to a real payment processor or EMV terminal.

---

**Author:** [Alby Kennady](https://github.com/albykennady)