# Go Microservices Project

This project is a collection of microservices written in Go. Each service has its own folder and is responsible for a specific task. The services communicate with each other using HTTP and RabbitMQ.

## Services

- **broker-service**: Handles requests and routes them to the correct service.
- **authentication**: Manages user authentication and authorization.
- **logger**: Logs events and errors from other services.
- **mail-service**: Sends emails for notifications and alerts.
- **listener-service**: Listens for events from RabbitMQ and processes them.
- **front-end**: The web front-end for users.

## How to Run

You can use the Makefile inside the `project` folder to manage and run your services easily.

### Common Commands

- `make up` - Start all Docker containers in the background.
- `make up_build` - Build all Go binaries, then start Docker containers (rebuilds images if needed).
- `make down` - Stop all Docker containers.
- `make build_broker` - Build the broker service binary for Linux.
- `make build_auth` - Build the authentication service binary for Linux.
- `make build_logger` - Build the logger service binary for Linux.
- `make build_mail` - Build the mail service binary for Linux.
- `make build_listener` - Build the listener service binary for Linux.
- `make build_front` - Build the front-end binary.
- `make start` - Build and start the front-end locally (without Docker).
- `make stop` - Stop the front-end running locally.

You can also run `docker-compose up` manually.

## Folder Structure

- Each service has its own folder with its code and Dockerfile.
- The `project` folder contains the `docker-compose.yaml` file to run all services together.

## Requirements

- Go
- RabbitMQ
- PostgreSQL
- MongoDB
- Docker

## Notes

- Services use RabbitMQ for messaging.
- MongoDB and PostgreSQL are used for data storage.
- You can add more services as needed.

