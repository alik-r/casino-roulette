# Casino Roulette 🎰
![Casino Roulette](./frontend/images/devil.png)

**Casino Roulette** is a **Cuphead-style** web app that simulates an roulette game where users bet their souls.
Register, deposit souls, place bets, and climb the leaderboard 😈.

## Table of Contents

- [Features](#features)
- [Technologies](#technologies)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Running with Docker Compose](#running-with-docker-compose)
  - [Running Manually](#running-manually)
- [API Endpoints](#api-endpoints)
- [Running Tests](#running-tests)
- [Contributing](#contributing)
- [License](#license)

## Features

- **User Authentication:** Registration and login using JWT.
- **Soul Deposits:** Users can deposit souls to their account.
- **Dynamic Betting:** Place bets on colors, even/odd, high/low, or specific numbers.
- **Leaderboard:** View top players based on soul balance.
- **Interactive UI:** Frontend with HTMX and Alpine.js.
- **Cuphead Aesthetic:** Unique style inspired by Cuphead.
- **Containerized Deployment:** Easy setup and deployment using Docker Compose.

## Technologies

- **Backend:** [Golang](https://golang.org/)
- **Frontend:** [HTMX](https://htmx.org/), [Alpine.js](https://alpinejs.dev/)
- **Database:** [PostgreSQL](https://www.postgresql.org/)
- **Containerization:** [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/)
- **Authentication:** JWT (JSON Web Tokens)
- **Web Framework:** [Chi](https://github.com/go-chi/chi)

## Project Structure

```
casino-roulette
├── backend
│   ├── cmd
│   │   └── main.go
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── pkg
│       ├── api
│       ├── auth
│       ├── db
│       ├── middleware
│       ├── models
│       └── roulette
├── docker-compose.yml
├── frontend
│   ├── Dockerfile
│   ├── fonts
│   │   ├── CupheadFelix-Regular-merged.ttf
│   │   ├── CupheadHenriette-A-merged.ttf
│   │   ├── CupheadMemphis-Medium-merged.ttf
│   │   ├── CupheadPoster-Regular.ttf
│   │   ├── CupheadVogue-Bold-merged.ttf
│   │   └── CupheadVogue-ExtraBold-merged.ttf
│   ├── images
│   │   ├── avatars
│   │   ├── devil.png
│   │   ├── leaderboard_icon.png
│   │   ├── logout_icon.png
│   │   ├── money_icon.png
│   │   ├── roulette_arrow.png
│   │   ├── roulette_circle.png
│   │   ├── roullete_icon.png
│   │   ├── skull_icon.png
│   │   ├── soul_icon.png
│   │   └── user_icon.png
│   ├── index.html
│   ├── leaderboard.html
│   ├── nginx.conf
│   ├── profile.html
│   ├── register.html
│   ├── roulette.html
│   ├── scripts
│   │   ├── avatar.js
│   │   └── script.js
│   ├── style.css
│   └── styles
│       ├── common.css
│       ├── leaderboard.css
│       ├── profile.css
│       ├── register.css
│       └── roulette.css
└── README.md
```

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your machine.
- [Docker Compose](https://docs.docker.com/compose/install/) installed.

### Installation

1. **Clone the repository:**

   ```sh
   git clone https://github.com/alik-r/casino-roulette.git
   cd casino-roulette
   ```

### Running with Docker Compose

The project can be fully launched using Docker Compose, which sets up both the backend and frontend along with the PostgreSQL database.

1. **Start the services:**

   ```sh
   docker-compose up --build
   ```

2. **Access the application:**

   - **Frontend:** [http://localhost:81](http://localhost:81)
   - **Backend API:** [http://localhost:8080](http://localhost:8080)

### Running Manually

If you prefer to run the backend and frontend separately without Docker Compose, follow these steps:

#### Backend

1. **Navigate to the backend directory:**

   ```sh
   cd backend
   ```

2. **Install dependencies:**

   ```sh
   go mod download
   ```

3. **Set Environment Variables:**

   Create a `.env` file or set the following environment variables:

   - `DATABASE_URL=postgres://gambler:ilovegambling@localhost:5433/casino?sslmode=disable`
   - `JWT_SECRET=super-secret-jwt-key`

4. **Run the server:**

   ```sh
   go run cmd/main.go
   ```

   The backend server will be running on [http://localhost:8080](http://localhost:8080).

#### Frontend

1. **Navigate to the frontend directory:**

   ```sh
   cd frontend
   ```

2. **Serve Static Pages:**

   You can use any static file server. For example, using `nginx`:

   ```sh
   docker build -t casino-roulette-frontend .
   docker run -d -p 81:80 casino-roulette-frontend
   ```

   Access the frontend at [http://localhost:81](http://localhost:81).

## API Endpoints

### Public Endpoints

- **Healthcheck**

  - **URL:** `/api/healthcheck`
  - **Method:** `GET`
  - **Response:**
    - **Status:** `200 OK`
    - **Body:** `"OK"`

- **Register**

  - **URL:** `/api/register`
  - **Method:** `POST`
  - **Request Body:**

    ```json
    {
      "username": "string",
      "email": "string",
      "password": "string",
      "avatar": "string" // Optional
    }
    ```

  - **Response:**

    ```json
    {
      "token": "string",
      "message": "Registered successfully"
    }
    ```

- **Login**

  - **URL:** `/api/login`
  - **Method:** `POST`
  - **Request Body:**

    ```json
    {
      "email": "string",
      "password": "string"
    }
    ```

  - **Response:**

    ```json
    {
      "token": "string"
    }
    ```

### Protected Endpoints

*All protected endpoints require the `Authorization` header with a valid JWT token:*

```
Authorization: Bearer <token>
```

- **Get User Information**

  - **URL:** `/api/user`
  - **Method:** `GET`
  - **Response:**

    ```json
    {
      "username": "string",
      "email": "string",
      "avatar": "string",
      "balance": 1000
    }
    ```

- **Deposit Souls**

  - **URL:** `/api/user/deposit`
  - **Method:** `POST`
  - **Request Body:**

    ```json
    {
      "amount": 100
    }
    ```

  - **Response:**

    ```json
    {
      "id": 1,
      "username": "string",
      "email": "string",
      "avatar": "string",
      "balance": 1100
    }
    ```

- **Place a Bet**

  - **URL:** `/api/roulette`
  - **Method:** `POST`
  - **Request Body:**

    ```json
    {
      "bet_amount": 50,
      "bet_type": "color", // Options: "color", "evenodd", "highlow", "number"
      "bet_value": "red" // Depending on bet_type, could be string or number
    }
    ```

  - **Response:**

    ```json
    {
      "balance": 1050,
      "bet_amount": 50,
      "bet_type": "color",
      "bet_value": "red",
      "payout": 50,
      "result": "win",
      "result_color": "red",
      "result_number": 32
    }
    ```

- **Get Leaderboard**

  - **URL:** `/api/leaderboard`
  - **Method:** `GET`
  - **Response:**

    ```json
    [
      {
        "username": "player1",
        "balance": 5000,
        "avatar": "images/avatars/avatar1.png"
      },
      {
        "username": "player2",
        "balance": 4500,
        "avatar": "images/avatars/avatar2.png"
      }
      // Top 10 players
    ]
    ```

## Running Tests

To ensure the integrity of the backend functionalities, run the tests using the following command:

```sh
cd backend
go test ./...
```