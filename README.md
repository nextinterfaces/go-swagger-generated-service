
# Sample API Service

This is a sample Golang HTTP API server generated with Swagger. The server is designed to provide a RESTful API based on the specifications in `swagger.yaml`.

## Project Structure

- `cmd/sample-api-server`: Main entry point for the API server.
- `restapi`: Contains the server logic, handler definitions, and middleware.
- `swagger.yaml`: Swagger specification file defining API endpoints, models, and documentation.

## Prerequisites

- **Golang**: Ensure Go is installed (version 1.17 or newer).
- **Swagger**: Install `go-swagger` to generate and validate Swagger documentation.

  ```bash
  go install github.com/go-swagger/go-swagger/cmd/swagger@latest
  ```

## Setup Instructions

1. **Clone the Repository** (if not done already)

    ```bash
    git clone <repository_url>
    cd go-swagger-generated-service1
    ```

2. **Install Dependencies**

    This project uses Go modules to manage dependencies. Run:

    ```bash
    go mod download
    ```

3. **Generate Code (Optional)**

    If modifications are made to `swagger.yaml`, you may need to regenerate server files:

    ```bash
    swagger generate server -f swagger.yaml -A sample-api
    ```

## Building and Running the Server

1. **Build the Server**

    ```bash
    go build -o sample-api-server ./cmd/sample-api-server
    ```

2. **Run the Server**

    Start the server, specifying the port (default is `8080`):

    ```bash
    ./sample-api-server --port=8080
    ```

3. **Verify the Server**

    You can test the API using `curl` or a browser.

    ```bash
    curl -X GET http://localhost:8080/api/v1/hello
    ```

## Swagger Documentation

The Swagger UI for the API documentation will be available at `http://localhost:8080/swagger`.

## Project Configuration

- **go.mod**: Manages project dependencies.
- **go-swagger**: Handles Swagger code generation.

## Additional Notes

For further configuration options, run:

```bash
./sample-api-server --help
```

This will display available flags, including options for setting host and port.
