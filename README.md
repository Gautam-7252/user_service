# Golang gRPC User Service

This repository contains a Golang gRPC service for managing user details, including functionalities to fetch, list, and search user details.

## Table of Contents

- [Golang gRPC User Service](#golang-grpc-user-service)
  - [Table of Contents](#table-of-contents)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
  - [Building and Running the Application](#building-and-running-the-application)
  - [Accessing the gRPC Service Endpoints](#accessing-the-grpc-service-endpoints)
  - [Configuration Details](#configuration-details)
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

1. Clone the repository:

   ```sh
   git clone https://github.com/yourusername/user_service.git
   cd user_service
