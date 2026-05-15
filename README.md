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



