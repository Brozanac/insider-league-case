# Insider League Simulation Case

Backend simulation of a football league system developed with Go for the Insider Development Intern Hiring Day case.

The application simulates a 4-team Premier League-style league, generates fixtures, plays matches, calculates standings, estimates championship probabilities, and allows match result editing through REST API endpoints.

---

## Tech Stack

- Go
- Gin
- GORM
- SQLite
- Docker
- REST API

---

## League Format

The system follows the original case requirement with a 4-team Premier League-style simulation.

The league contains:

- 4 teams
- 6 weeks
- 2 matches per week
- 12 total matches
- Each team plays 6 matches
- Premier League-style points and standings rules

Standings are calculated based on:

1. Points
2. Goal difference
3. Goals scored

---

## Features

- Initialize league teams and fixtures
- Generate double round-robin fixtures
- Simulate selected weeks
- Simulate all remaining league matches
- Calculate league standings dynamically
- List all matches and results
- Estimate championship probabilities
- Edit match results manually
- Reset league data during initialization

---

## Extra Features

In addition to the core requirements, this project includes:

- Automatic league play until the end of the season
- Editable match results
- Dynamic standings recalculation
- Monte Carlo championship prediction system
- Docker support
- SQL schema and example queries

---

## Architecture

This project uses a layered backend architecture:

```txt
HTTP Handlers
↓
Services
↓
Repositories
↓
Database