# Casino Roulette

Casino Roulette is a web application that simulates a roulette game. Users can register, deposit money, place bets, and view the leaderboard.

## Planned Project Structure
```
/casino-roulette
├── backend
│   ├── cmd
│   │   └── main.go
│   ├── pkg
│   │   ├── api
│   │   │   └── handlers.go
│   │   ├── models
│   │   │   └── models.go
│   │   ├── roulette
│   │   │   └── roulette.go
│   │   └── db
│   │       └── db.go
│   ├── config
│   │   └── config.go
│   ├── Dockerfile
│   └── go.mod
├── frontend
│   ├── index.html
│   ├── style.css
│   └── app.js
│   └── Dockerfile
├── docker-compose.yml
├── .github
│   └── workflows
│       └── ci.yml
└── README.md
```


## Getting Started

### Prerequisites

- Go 1.22 or later
- SQLite

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/alik-r/casino-roulette.git
    cd casino-roulette/backend
    ```

2. Install dependencies:

    ```sh
    go mod install
    ```

## Running the Application

To start the server, run:

```sh
go run cmd/main.go
```

The server will be running on `http://localhost:8080`.

## API Endpoints

### Login 
- **URL:** `/api/user/login`
- **Method:** `POST`
- **Request Body:**

    ```json
    {
        "username": "string",
        "password": "string"
    }
    ```

- **Response:**

    ```json
    {
        "token": "string",
        "is_register": "boolean"
    }
    ```

### Deposit

- **URL:** `/api/user/deposit`
- **Method:** `POST`
- **Request Body:**

    ```json
    {
        "amount": "int"
    }
    ```

- **Response:**

    ```json
    {
        "id": "int",
        "username": "string",
        "balance": "int"
    }
    ```

### Place Bet

- **URL:** `/api/roulette/bet`
- **Method:** `POST`
- **Request Body:**

    ```json
    {
        "bet_amount": "int",
        "bet_color": "{red|black|green}"
    }
    ```

- **Response:**

    ```json
    {
        "username": "string",
        "balance": "int",
        "bet_amount": "int",
        "bet_color": "{red|black|green}",
        "result": "{win|lose}",
        "result_color": "{red|black|green}",
        "username": "string"
    }
    ```

### Get Leaderboard

- **URL:** `/api/leaderboard`
- **Method:** `GET`
- **Response:**

    ```json
    [
        {
            "id": "int",
            "username": "string",
            "balance": "int"
        }
    ]
    ```

## Running Tests

To run the tests, use:

```sh
go test ./...
```
