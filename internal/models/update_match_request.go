package models

type UpdateMatchRequest struct {
	HomeGoals int `json:"home_goals"`
	AwayGoals int `json:"away_goals"`
}
