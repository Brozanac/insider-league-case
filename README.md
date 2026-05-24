# Insider League Case

This is my solution for the Insider Development Intern Hiring Day case.

The project simulates a small football league with 4 teams. It generates fixtures, plays matches week by week, calculates the league table, estimates championship probabilities, and allows match results to be edited manually.

I built the backend with Go and added a small React dashboard to make the API easier to test and present.

---

## What the project does

- Creates a 4-team league
- Generates a double round-robin fixture list
- Simulates a selected week
- Simulates all remaining matches
- Calculates the league table dynamically
- Estimates championship probabilities with Monte Carlo simulation
- Allows match results to be edited
- Updates standings and predictions after edited results
- Provides a simple React dashboard
- Can be run locally or with Docker

---

## Tech stack

### Backend

- Go
- Gin
- GORM
- SQLite

### Frontend

- React
- Vite
- Axios

### Other

- Docker
- Docker Compose

---

## Project structure

```txt
insider-league-case/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   ├── internal/
│   │   ├── database/
│   │   ├── handlers/
│   │   ├── models/
│   │   ├── repositories/
│   │   └── services/
│   ├── docs/
│   │   ├── schema.sql
│   │   └── example_queries.sql
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   │   ├── api/
│   │   ├── App.jsx
│   │   ├── App.css
│   │   └── main.jsx
│   ├── Dockerfile
│   ├── package.json
│   └── vite.config.js
│
├── docker-compose.yml
├── README.md
└── .gitignore


## League Format

The system follows the original case requirement with a 4-team Premier League-style simulation.

The league contains:

* 4 teams
* 6 weeks
* 2 matches per week
* 12 total matches
* Each team plays 6 matches
* Premier League-style points and standings rules

Standings are sorted by:

1. Points
2. Goal difference
3. Goals scored

## Features
* Initialize league teams and fixtures
* Generate double round-robin fixtures
* Simulate a selected week
* Simulate all remaining league matches
* Calculate standings dynamically
* List all fixtures and match results
* Estimate championship probabilities
* Edit match results manually
* Recalculate standings after edited results
* Reset league data during initialization
* Run backend and frontend with Docker Compose

## Extra Features

In addition to the core requirements, this project includes:

* Automatic league play until the end of the season
* Editable match results
* Dynamic standings recalculation
* Monte Carlo championship prediction system
* React frontend dashboard
* Docker support for both backend and frontend
* SQL schema and example queries

## Architecture

The backend uses a layered architecture:
```
HTTP Handlers
↓
Services
↓
Repositories
↓
Database
```

The project follows:

* Interface-based design
* Repository pattern
* Service layer pattern
* Struct composition
* Interface-Based Design

Interfaces are used to decouple business logic from concrete implementations.

For example, services depend on repository interfaces instead of directly depending on GORM or SQLite. This makes the project easier to test, maintain, and extend.

### Repository Pattern

Repositories handle database operations such as creating, fetching, updating, and deleting records.

This keeps persistence logic separate from business logic.

### Service Layer

Services contain the main business rules, including:

* League initialization
* Fixture generation
* Match simulation
* Standings calculation
* Championship prediction
* Editable match results

Handlers stay lightweight and only handle request parsing, validation, service calls, and JSON responses.

## API Endpoints

### Base URL:
```
http://localhost:8080
```

### Initialize League:

```
POST /league/init
```

Creates teams and fixtures.

Example response:

```
{
  "message": "League initialized successfully"
}
```

### Get League Table
GET /league/table

Returns current league standings.

Example response:
```
[
  {
    "team_id": 1,
    "team_name": "Chelsea",
    "played": 2,
    "won": 1,
    "drawn": 1,
    "lost": 0,
    "goals_for": 4,
    "goals_against": 2,
    "goal_difference": 2,
    "points": 4
  }
]
```

### Get All Matches

```
GET /matches
```

Returns all fixtures and match results.

Example response:
```
[
  {
    "id": 1,
    "week": 1,
    "home_team_id": 1,
    "away_team_id": 2,
    "home_goals": 2,
    "away_goals": 1,
    "played": true
  }
]
```

### Play Selected Week

```
POST /league/play/week/{week}
```

Valid week range: 1 to 6



Example:

```
POST /league/play/week/1
```

Example response:
```
{
  "message": "Week played successfully",
  "week": 1
}
```

### Play All Remaining Matches

```
POST /league/play/all
```

Simulates all unplayed matches until the league is completed.

Example response:

```
{
  "message": "All remaining matches played successfully"
}
```

### Get Championship Predictions

```
GET /league/predictions
```

Returns championship probability estimates for each team.

Example response:
```
[
  {
    "team_id": 3,
    "team_name": "Manchester City",
    "probability": 42.6
  },
  {
    "team_id": 4,
    "team_name": "Liverpool",
    "probability": 31.4
  }
]
```

### Update Match Result

```
PUT /matches/{id}
```

Updates the selected match result and marks it as played.

Example:

```
PUT /matches/1
```

Request body:

```
{
  "home_goals": 2,
  "away_goals": 1
}
```

Example response:
```
{
  "message": "Match result updated successfully"
}
```


The league table and predictions are recalculated dynamically based on the updated match data.


### Prediction System

The project includes a Monte Carlo-based championship prediction system.

The prediction engine works by cloning the current league state and simulating the remaining unplayed matches multiple times.

After each simulation:

* Remaining matches are simulated.
* The final league table is calculated.
*The champion is recorded.

The final probability is calculated as:

championship_probability = champion_count / total_simulations * 100

This allows the API to estimate each team’s chance of finishing first based on the current league table and remaining fixtures.

## Frontend Dashboard

A React frontend is included under the frontend/ directory.

The dashboard supports:

* Initializing the league
* Viewing the league table
* Viewing weekly fixtures
* Viewing all matches
* Playing a selected week
* Playing all remaining matches
* Editing match results
* Viewing championship predictions

Frontend URL:
```
http://localhost:5173
```

Backend URL:
```
http://localhost:8080
```

## How To Run

### Run Locally
1. Run Backend

From the project root:

```
cd backend
go mod tidy
go run cmd/api/main.go
```

The backend will be available at:
```
http://localhost:8080
```

2. Run Frontend

Open a second terminal from the project root:

```
cd frontend
npm install
npm run dev
```

The frontend will be available at:

```
http://localhost:5173
```

### Run with Docker

This project includes Docker support for both backend and frontend.

From the project root:
```
docker compose up --build
```

Backend:
```
http://localhost:8080
```

Frontend:
```
http://localhost:5173
```

Stop containers:
```
docker compose down
```

#### Docker Configuration
Root `docker-compose.yml`

```
services:
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8080:8080"

  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "5173:5173"
    depends_on:
      - backend
```

#### Backend Dockerfile

Located at:
```
backend/Dockerfile
```

#### Frontend Dockerfile

Located at:

```
frontend/Dockerfile
```

### SQL Schema and Queries

The project includes SQL files under:

```
backend/docs/
```

Files:
```
backend/docs/schema.sql
backend/docs/example_queries.sql
```

schema.sql contains database table definitions.

example_queries.sql contains useful SQL examples for:


* inspecting teams
* inspecting matches
* filtering played matches
* filtering unplayed matches
* updating match results
* resetting SQLite auto-increment values

#### Example SQL Schema
```
CREATE TABLE teams (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    strength INTEGER NOT NULL
);

CREATE TABLE matches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    week INTEGER NOT NULL,
    home_team_id INTEGER NOT NULL,
    away_team_id INTEGER NOT NULL,
    home_goals INTEGER NOT NULL DEFAULT 0,
    away_goals INTEGER NOT NULL DEFAULT 0,
    played BOOLEAN NOT NULL DEFAULT FALSE,

    FOREIGN KEY (home_team_id) REFERENCES teams(id),
    FOREIGN KEY (away_team_id) REFERENCES teams(id)
);

```
## Additional Notes

* The backend can be tested independently using Postman, Thunder Client, curl, or any REST client.
* The frontend is included as a visual demo layer.
* League standings are calculated dynamically from match results.
* Editing a match result automatically affects standings and predictions.
* SQLite is used for simple local development.