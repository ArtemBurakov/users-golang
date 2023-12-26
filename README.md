# Users API in Golang

This repository contains a simple CRUD (Create, Read, Update, Delete) API for managing users implemented in Golang. It utilizes the Gin framework for handling HTTP requests.

## Prerequisites

Before running the project, ensure that you have the following installed on your system:

- [Go](https://golang.org/dl/) (Golang)
- [Docker](https://www.docker.com/get-started)

## Getting Started

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/users-api-golang.git
   cd users-api-golang

2. Create a .env file in the project root and set the required environment variables
    ```bash
   PORT=8045
   GIN_MODE=release

3. Build and run the project using the provided Makefile commands
    ```bash
   make docker-all

4. Access the API at http://localhost:8045/api/users (assuming you set PORT=8045 in your .env file).