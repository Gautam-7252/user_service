# Golang gRPC User Service

This repository contains a Golang gRPC service for managing user details, including functionalities to fetch, list, and search user details.

## Table of Contents

- [Golang gRPC User Service](#golang-grpc-user-service)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Building and Running the Application](#building-and-running-the-application)
  - [Accessing the gRPC Service Endpoints](#accessing-the-grpc-service-endpoints)
  - [Running Tests](#running-tests)
  - [Testing with grpcurl](#testing-with-grpcurl)
  - [Dockerizing the Application](#dockerizing-the-application)
  - [Contributing](#contributing)
  - [License](#license)

## Prerequisites

- Go 1.22.4
- `protoc` (Protocol Buffers compiler)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins
- `grpcurl` (for manual testing)

## Setup

1. Install Go dependencies:

   ``` sh
   git clone https://github.com/yourusername/user_service.git
   cd user_service

2. Clone the repository:

   ``` sh
   go mod tidy

3. Install Protocol Buffers and gRPC plugins:

   ``` sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

4. Generate Go code from Protocol Buffers definitions:

   ``` sh
   protoc --go_out=. --go-grpc_out=. proto/user.proto

## Building and Running the Application

1. Build the application:

   ``` sh
   go build -o user_service main.go

2. Run the application:

   ``` sh
   ./user_service
