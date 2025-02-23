# Alluvial Takehome

This is a simple project built with Go and Docker to demonstrate various application deployment processes.

## Setup

To get started, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/avawing/alluvial-takehome.git
    ```
2. Navigate to the project directory:
    ```bash
    cd alluvial
    ```
3. Copy the `.env.example` to `.env` and configure any required environment variables:
    ```bash
    cp .env.example .env
    ```

## Dependencies

The project uses the following dependencies:
- **Go**: The main programming language for the application.
- **Docker**: For containerization.

## Running Locally

To build and run the application locally using Docker & Kubernetes:

1. Build the Docker image:
    ```bash
    docker compose build .
    ```
2. Run the application:
    ```bash
    docker compose push 
    ```
3. ```bash
    minikube start
   ```
4. ```bash
   kubectl config use-context minikube
   ```
5. ```bash
   kubectl apply -f k8s/deployment.yaml
   ```
5. ```bash
   kubectl port-forward svc/alluvial-app-service 8080:8080
   ```
To build the project using only docker-compose:
```bash
   docker-compose up -d
   ```

## Project Structure

- **handlers**: Contains HTTP request handlers.
- **k8s**: Kubernetes configurations for deployment.
- **models**: Data models and business logic.
- **repository**: Database interaction logic.
- **services**: External API/service interactions.
- **utils**: Helper utilities for the project.

## Additional Information

For any questions or issues, feel free to open an issue in the repository.
