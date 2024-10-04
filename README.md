# Golang Basic REST API

This project is a basic REST API built with Golang.

## Features

- Basic CRUD operations
- RESTful endpoints
- JSON responses
- Database & Migration system in SQLite

## Requirements

- Go 1.16 or higher

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/sherwin-77/golang-basic-rest-api.git
    ```
2. Navigate to the project directory:
    ```sh
    cd golang-basic-rest-api
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Migrating

All migrate command available in `cmd/migrate`

#### Creating Migration
```sh
go run cmd/migrate create create_any_tables
```
#### Apply Migration
```sh
go run cmd/migrate up
```

#### Rollback Migration
```sh
go run cmd/migrate down
```

DB path can be customized in `config.json`


## Usage

Make sure you already run migration
```sh
go run cmd/migrate up
```

1. Run the server:
    ```sh
    go run cmd/server
    ```
2. The API will be available at `http://localhost:8080`.

## Endpoints

- `GET /todos` - Retrieve all todos
- `GET /todos/{id}` - Retrieve a specific todo by ID
- `POST /todos` - Create a new todo
- `PATCH /todos/{id}` - Update an existing todo by ID
- `DELETE /todos/{id}` - Delete an todo by ID


## Building

SoonTM