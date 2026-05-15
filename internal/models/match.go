package models

type Match struct {
	ID   uint `json:"id" gorm:"primaryKey"`
	Week int  `json:"week"`

	HomeTeamID uint `json:"home_team_id"`
	AwayTeamID uint `json:"away_team_id"`

	HomeGoals int  `json:"home_goals"`
	AwayGoals int  `json:"away_goals"`
	Played    bool `json:"played"`
}
