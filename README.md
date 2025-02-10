# Password Validator API

This REST API was built in Go using the [Chi](https://go-chi.io/) framework. This is a simple microservice that provides a RESTful API endpoint that can validate passwords based on a set of predefined rules.

## Features

- RESTful API endpoint for password validation (`/v1/users/password-validation`).
- Modular rule-based validation system.
- Easy to extend and create new rules for validation.

### Sequence Diagram

<!-- TODO: Add Sequence Diagram -->

## Available validators

The application provides validators for the following rules:

- MinLengthRule
- DigitRule
- LowercaseRule
- UppercaseRule
- SpecialCharRule
- NoRepeatedCharsRule

## Project Structure

```
password-validator-api
│   go.mod    
│   go.sum    
│   README.md
│   request.http    
│
└───cmd
│   └───api
│       └───main
│           │   main.go
│   
└───internal
│   └───api
│       └───handlers
│           │   password_handler_test.go
│           │   password_handler.go
│   └───api
│       └───domain
│           └───password
│               │   rules.go
│               │   validator_test.go
│               │   validator.go
```

## API Usage

### Validate Password

```bash
POST /v1/users/validate-password
Content-Type: application/json

{
    "password": "YourPassword123!"
}
```

### Response

```json
{
    "valid": true,
    "errors": []
}
```

## Getting Started

### Requirements

- - Go 1.23.0 or higher
- _(optional)_ Docker

### Installation

#### If you're not using Docker

1. Clone the repository

```bash
git clone https://github.om/vitorvasc/password-validator-api.git
```

2. Run the application

```bash
go run cmd/api/main/main.go
```

The server will start on port 8080 by default. 

#### If you're using Docker

```bash
docker-compose up --build
```

## Testing

* Run the tests:
```bash
go test -v ./...
```

* Run the tests with coverage:

```bash
go test -v ./... -cover
```

## Contributing

Feel free to submit issues, fork the repository and create pull requests for any improvements. 
