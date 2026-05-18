# Backend simulation of a football league system developed with Go.

# Features : 

- League simulation
- Match generation
- League standings
- Championship prediction
- REST API

# Tech Stack : 
- Go
- Gin
- GORM
- SQLite

Setup : 
## Run

go run cmd/api/main.go


# API Endpoints : 

POST /league/init
GET  /league/table

# Architecture

This project was designed using a layered backend architecture to ensure
maintainability, scalability, readability, and separation of concerns.

The application follows an interface-driven design approach together with
Repository and Service patterns.

---

## High-Level Architecture

Client Request
↓
HTTP Handlers
↓
Services
↓
Repositories
↓
Database

Each layer has a single responsibility and communicates only with the layer
directly below it.

---

# Interface-Based Design

The project heavily uses interfaces in order to decouple implementation details
from business logic.

Example:

```go
type TeamRepository interface {
	Create(team *models.Team) error
	FindAll() ([]models.Team, error)
} ```go


## Prediction System

The project includes a Monte Carlo-based championship prediction system.

The prediction engine works by cloning the current league state and simulating
the remaining unplayed matches multiple times. After each simulation, the final
league table is calculated and the champion is recorded.

The final probability is calculated as:

winner_count / total_simulations * 100

This allows the API to estimate each team's chance of finishing first based on
the current table and remaining fixtures.


## Run with Docker

Build and start the application:

```bash
docker compose up --build

The league contains:

- 4 teams
- 6 weeks
- 2 matches per week
- 12 total matches
- Premier League-style points and standings rules

## SQL Schema and Queries

The project includes SQL files under the `docs/` directory.

```txt
docs/schema.sql
docs/example_queries.sql