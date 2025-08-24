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

## Deployment Options

### Docker Compose

Navigate to the `project` directory and run:

```bash
# Start everything
make up_build

# Or just start (without rebuild)
make up

# Stop all services
make down
```

### Kubernetes Deployment

Deploy to Kubernetes cluster:

```bash
# Deploy all services
kubectl apply -f project/k8s/

# Check deployment status
kubectl get pods

# Access services via ingress
# Frontend: http://front-end.info
# Broker API: http://broker-service.info
```

**Note**: Ensure NGINX Ingress Controller is installed and `/etc/hosts` is configured for local testing:
```
127.0.0.1 front-end.info
127.0.0.1 broker-service.info
```

## Tech Stack

**Backend**: Go, gRPC, REST APIs  
**Databases**: PostgreSQL, MongoDB  
**Messaging**: RabbitMQ (AMQP)  
**Email**: SMTP, MailHog (development)  
**Deployment**: Docker Compose, Kubernetes  
**Orchestration**: NGINX Ingress, Resource Management  
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

## Kubernetes Features

- **Resource Management**: CPU/Memory requests and limits configured
- **Horizontal Scaling**: Ready for replica scaling
- **Service Discovery**: Internal service communication via DNS
- **Ingress Routing**: External access through NGINX ingress
- **Health Checks**: Kubernetes-native health monitoring
- **Rolling Updates**: Zero-downtime deployments

## Project Structure

```
├── authentication/     # User management service
├── broker-service/     # API gateway
├── logger/            # Logging service with gRPC
├── mail-service/      # Email service
├── listener-service/  # Event consumer
├── front-end/         # Web interface
└── project/           # Orchestration & deployment
    ├── docker-compose.yaml
    ├── k8s/           # Kubernetes manifests
    │   ├── ingress.yaml      # NGINX ingress controller
    │   ├── broker.yaml       # Broker deployment & service
    │   ├── authentication.yaml
    │   ├── logger.yaml
    │   ├── mail.yaml
    │   ├── listener.yaml
    │   ├── front-end.yaml
    │   ├── mongo.yaml        # MongoDB deployment
    │   ├── rabbit.yaml       # RabbitMQ deployment
    │   └── mailhog.yaml      # MailHog for email testing
    └── Makefile
```

## Prerequisites

**For Docker Deployment:**
- Go 1.19+
- Docker & Docker Compose
- Make

**For Kubernetes Deployment:**
- kubectl
- Kubernetes cluster (local or cloud)
- NGINX Ingress Controller

---
*[Habibur Rahman](https://habib.im)*

