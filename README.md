# Password Validator API

This REST API was built in Go using the [Chi](https://go-chi.io/) framework. This is a simple microservice that provides a RESTful API endpoint that can validate passwords based on a set of predefined rules.

## Table of Contents

- [Password Validator API](#password-validator-api)
    - [Features](#features)
        - [Sequence Diagram](#sequence-diagram)
    - [Available validators](#available-validators)
    - [Project structure](#project-structure)
    - [API Usage](#api-usage)
        - [Validate password](#validate-password)
        - [Response](#response)
            - [Valid password](#valid-password)
            - [Invalid password](#invalid-password)
    - [Getting Started](#getting-started)
        - [Requirements](#requirements)
        - [Installation](#installation)
            - [If you're not using Docker](#if-youre-not-using-docker)
            - [If you're using Docker](#if-youre-using-docker)
    - [Testing](#testing)
    - [Contributing](#contributing)

## Features

- RESTful API endpoint for password validation (`/v1/users/password-validation`).
- Modular rule-based validation system.
- Easy to extend and create new rules for validation.

### Sequence Diagram

<!-- TODO: Add Sequence Diagram -->

## Available validators

The application provides validators for the following rules:

- [DigitRule](https://github.com/vitorvasc/password-validator-api/blob/main/internal/domain/rule/digit_rule.go)
- [LowercaseRule](https://github.com/vitorvasc/password-validator-api/blob/main/internal/domain/rule/lowercase_rule.go)
- [MinLengthRule](https://github.com/vitorvasc/password-validator-api/blob/main/internal/domain/rule/min_length_rule.go)
- [NoRepeatedCharsRule](https://github.com/vitorvasc/password-validator-api/blob/main/internal/domain/rule/no_repeated_chars_rule.go)
- [SpecialCharRule](https://github.com/vitorvasc/password-validator-api/blob/main/internal/domain/rule/special_char_rule.go)
- [UppercaseRule](https://github.com/vitorvasc/password-validator-api/blob/main/internal/domain/rule/uppercase_rule.go)

## Project structure

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
│           │   password_handler.go
|
│       └───domain
│           └───rule
│               │   digit_rule.go
│               │   lowercase_rule.go
│               │   min_length_rule.go
│               │   no_repeated_chars_rule.go
│               │   special_char_rule.go
│               │   uppercase_rule.go
|
│           └───validator
│               │   validator_option.go
│               │   validator.go
|
```

## API Usage

### Validate password

```bash
POST /v1/users/validate-password
Content-Type: application/json

{
    "password": "YourPassword123!"
}
```

**cURL**:

```
curl --location 'http://localhost:8080/v1/users/validate-password' \
    --header 'Content-Type: application/json' \
    --data '{
        "password": "YourPassword123!"    
    }'
```

### Response

#### Valid password

```json
{
    "valid": true,
    "errors": []
}
```

#### Invalid password

```json
{
    "valid": false,
    "errors": [
        "must be at least 9 characters long"
    ]
}
```

## Getting Started

### Requirements

- Go 1.23.0 or higher
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
