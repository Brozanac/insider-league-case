-- Get all teams
SELECT *
FROM teams;

-- Get all matches ordered by week
SELECT *
FROM matches
ORDER BY week ASC, id ASC;

-- Get matches for a selected week
SELECT *
FROM matches
WHERE week = 1
ORDER BY id ASC;

-- Get played matches
SELECT *
FROM matches
WHERE played = true
ORDER BY week ASC, id ASC;

-- Get unplayed matches
SELECT *
FROM matches
WHERE played = false
ORDER BY week ASC, id ASC;

-- Get a single match by ID
SELECT *
FROM matches
WHERE id = 1;

-- Update a match result manually
UPDATE matches
SET home_goals = 3,
    away_goals = 1,
    played = true
WHERE id = 1;

-- Delete all matches
DELETE FROM matches;

-- Delete all teams
DELETE FROM teams;

-- Reset SQLite auto-increment for matches
DELETE FROM sqlite_sequence
WHERE name = 'matches';

-- Reset SQLite auto-increment for teams
DELETE FROM sqlite_sequence
WHERE name = 'teams';

-- Basic home-side match contribution
SELECT
    t.id AS team_id,
    t.name AS team_name,
    COUNT(m.id) AS played,
    SUM(
        CASE
            WHEN m.home_goals > m.away_goals THEN 3
            WHEN m.home_goals = m.away_goals THEN 1
            ELSE 0
        END
    ) AS points
FROM teams t
JOIN matches m ON m.home_team_id = t.id
WHERE m.played = true
GROUP BY t.id, t.name;