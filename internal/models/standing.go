package models

type Standing struct {
	TeamID       uint   `json:"team_id"`
	TeamName     string `json:"team_name"`
	Played       int    `json:"played"`
	Won          int    `json:"won"`
	Drawn        int    `json:"drawn"`
	Lost         int    `json:"lost"`
	GoalsFor     int    `json:"goals_for"`
	GoalsAgainst int    `json:"goals_against"`
	GoalDiff     int    `json:"goal_difference"`
	Points       int    `json:"points"`
}
