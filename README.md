# Golang gRPC User Service

This repository contains a Golang gRPC service for managing user details, including functionalities to fetch, list, and search user details.

## Table of Contents

- [Golang gRPC User Service](#golang-grpc-user-service)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Running the Application and Test file](#running-the-application-and-test-file)
  - [Testing with Postman](#testing-with-postman)
  - [Dockerizing the Application](#dockerizing-the-application)

## Prerequisites

- Go 1.22.4
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins
- `grpcurl` (for manual testing)

## Setup

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/user_service.git
   cd user_service
2. Install Go dependencies:

   ```
   go mod tidy
3. Generate Go code from Protocol Buffers definitions:

    ```
    protoc --go_out=. --go-grpc_out=. proto/user.proto
## Running the Application and Test file

1. Run the gRPC server:
   
    ```
    go run main.go
2. Running test file:

   ```
   go test ./...
## Testing with Postman

### Setting Up Postman for gRPC

1. <b>Open Postman</b> and create a new gRPC request.
2. Enter the gRPC server URL: localhost:50051.
3. Import the Proto File:
   - Click on `Import` in Postman.
   - Import the `proto/user.proto` file.
  
### Testing Methods

1. GetUser
   - Method: `GetUser`
   - Request Body:
     ```
     {
      "id": 1
     }
2. ListUsers
   - Method: `ListUsers`
   - Request Body:
     ```
     {
     "ids": [1, 2, 3]
     }
3. SearchUsers
   - Method: `SearchUsers`
   - Request Body:
     ```
     {
      "city": "LA"
     }
