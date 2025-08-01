# Go Microservices

A modern microservices architecture built with Go, featuring event-driven communication and containerized deployment.

## Architecture

| Service | Port | Purpose | Technology |
|---------|------|---------|------------|
| **Broker** | 8080 | API Gateway & Request Routing | HTTP REST APIs |
| **Authentication** | 8081 | User Auth & Authorization | PostgreSQL, JWT |
| **Logger** | - | Centralized Logging via gRPC | MongoDB, gRPC |
| **Mail** | - | Email Notifications | SMTP, Go templates |
| **Listener** | - | RabbitMQ Event Processing | AMQP, Event-driven |
| **Frontend** | - | Web Interface | Go HTML templates |

### Service Details

- **Broker Service**: Entry point for all requests, routes to appropriate microservices
- **Authentication Service**: Manages user registration, login, and JWT token validation
- **Logger Service**: Receives log data via gRPC calls and stores in MongoDB
- **Mail Service**: Sends emails using configurable SMTP settings with HTML/text templates
- **Listener Service**: Consumes events from RabbitMQ queue for asynchronous processing
- **Frontend Service**: Serves web pages and handles user interactions

## Quick Start

Navigate to the `project` directory and run:

```bash
# Start everything
make up_build

# Or just start (without rebuild)
make up

# Stop all services
make down
```

## Tech Stack

**Backend**: Go, gRPC, REST APIs  
**Databases**: PostgreSQL, MongoDB  
**Messaging**: RabbitMQ (AMQP)  
**Email**: SMTP, MailHog (development)  
**Deployment**: Docker, Docker Compose  
**Build**: Make, Go modules

## Development

```bash
# Build individual services
make build_broker
make build_auth
make build_logger
make build_mail
make build_listener

# Run frontend locally
make start
make stop
```

## Project Structure

```
├── authentication/     # User management service
├── broker-service/     # API gateway
├── logger/            # Logging service with gRPC
├── mail-service/      # Email service
├── listener-service/  # Event consumer
├── front-end/         # Web interface
└── project/           # Docker orchestration
    ├── docker-compose.yaml
    └── Makefile
```

## Prerequisites

- Go 1.19+
- Docker & Docker Compose
- Make

---
*[Habibur Rahman](https://habib.im)*

