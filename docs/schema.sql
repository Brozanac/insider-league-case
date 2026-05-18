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