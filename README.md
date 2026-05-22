# Insider League Simulation Case

A full-stack football league simulation project developed for the Insider Development Intern Hiring Day case.

The project simulates a 4-team Premier League-style league, generates fixtures, plays matches, calculates standings, estimates championship probabilities, allows manual match result editing, and provides a React dashboard for visual interaction.

The original case requires a Go backend, interface-based design, struct composition, API-accessible endpoints, SQL schema/queries, and setup/deployment documentation. This project also includes bonus features such as a React frontend, editable match results, Monte Carlo predictions, and Docker support.

---

## Tech Stack

### Backend

- Go
- Gin
- GORM
- SQLite
- Docker

### Frontend

- React
- Vite
- Axios
- CSS

---

## Project Structure

```txt
insider-league-case/
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go
│   │
│   ├── internal/
│   │   ├── database/
│   │   │   └── database.go
│   │   │
│   │   ├── handlers/
│   │   │   └── league_handler.go
│   │   │
│   │   ├── models/
│   │   │   ├── match.go
│   │   │   ├── prediction.go
│   │   │   ├── standing.go
│   │   │   ├── team.go
│   │   │   └── update_match_request.go
│   │   │
│   │   ├── repositories/
│   │   │   ├── match_repository.go
│   │   │   └── team_repository.go
│   │   │
│   │   └── services/
│   │       ├── fixture_service.go
│   │       ├── league_service.go
│   │       ├── match_simulator.go
│   │       ├── prediction_service.go
│   │       └── standing_service.go
│   │
│   ├── docs/
│   │   ├── schema.sql
│   │   └── example_queries.sql
│   │
│   ├── Dockerfile
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── public/
│   ├── src/
│   │   ├── api/
│   │   │   └── leagueApi.js
│   │   ├── App.jsx
│   │   ├── App.css
│   │   └── main.jsx
│   │
│   ├── Dockerfile
│   ├── package.json
│   ├── package-lock.json
│   └── vite.config.js
│
├── docker-compose.yml
├── README.md
├── .gitignore
└── LICENSE
```


## League Format

The system follows the original case requirement with a 4-team Premier League-style simulation.

The league contains:

4 teams
6 weeks
2 matches per week
12 total matches
Each team plays 6 matches
Premier League-style points and standings rules

Standings are sorted by:

Points
Goal difference
Goals scored

## Features
Initialize league teams and fixtures
Generate double round-robin fixtures
Simulate a selected week
Simulate all remaining league matches
Calculate standings dynamically
List all fixtures and match results
Estimate championship probabilities
Edit match results manually
Recalculate standings after edited results
Reset league data during initialization
Run backend and frontend with Docker Compose

## Extra Features

In addition to the core requirements, this project includes:

Automatic league play until the end of the season
Editable match results
Dynamic standings recalculation
Monte Carlo championship prediction system
React frontend dashboard
Docker support for both backend and frontend
SQL schema and example queries

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

Interface-based design
Repository pattern
Service layer pattern
Struct composition
Interface-Based Design

Interfaces are used to decouple business logic from concrete implementations.

For example, services depend on repository interfaces instead of directly depending on GORM or SQLite. This makes the project easier to test, maintain, and extend.

### Repository Pattern

Repositories handle database operations such as creating, fetching, updating, and deleting records.

This keeps persistence logic separate from business logic.

### Service Layer

Services contain the main business rules, including:

league initialization
fixture generation
match simulation
standings calculation
championship prediction
editable match results

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