# Insider League Simulation Case

A full-stack football league simulation project developed for the Insider Development Intern Hiring Day case.

The project simulates a 4-team Premier League-style league, generates fixtures, plays matches, calculates standings, estimates championship probabilities, allows editable match results, and provides a React dashboard for visual interaction.

The original case requires a Go backend, interface-based design, struct composition, API-accessible endpoints, SQL schema/queries, and optionally deployment/setup documentation. This project includes those requirements plus a frontend dashboard and Docker support. :contentReference[oaicite:0]{index=0}

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
│   ├── public/
│   ├── src/
│   │   ├── api/
│   │   │   └── leagueApi.js
│   │   ├── App.jsx
│   │   ├── App.css
│   │   └── main.jsx
│   ├── Dockerfile
│   ├── package.json
│   ├── package-lock.json
│   └── vite.config.js
│
├── docker-compose.yml
├── README.md
├── .gitignore
└── LICENSE