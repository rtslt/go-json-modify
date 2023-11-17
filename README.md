# Go API with CORS Middleware

This project is a simple Go API that uses the CORS middleware to handle Cross-Origin Resource Sharing.

## Getting Started

To run this project, you will need to have Go installed on your machine.

## Code Overview

The `main.go` file contains the main logic of the application.

### Importing Required Packages

The following packages are imported:

- `fmt` and `net/http` from the standard library.
- `github.com/go-chi/chi` for routing.
- `github.com/go-chi/chi/middleware` for middleware support.
- `github.com/rs/cors` for handling CORS.

### CORS Middleware

The CORS middleware is set up with the following options:

- All origins are allowed.
- The allowed methods are GET, POST, PUT, DELETE, and OPTIONS.
- The allowed headers are Accept, Authorization, Content-Type, and X-CSRF-Token.
- The exposed headers are Link.
- Credentials are allowed.
- The maximum age for CORS preflight requests is 300.

### Routes

There are two routes defined:

- A GET route at the path `/`, which responds with "Hello World!".
- A GET route at the path `/pokemon/{pokemon}`, which makes a request to the PokeAPI for the specified Pokemon and prints the response.

## Running the Server

To run the server, use the command `go run main.go`. The server will start on port 3000.