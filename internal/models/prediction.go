package models

type Prediction struct {
	TeamID      uint    `json:"team_id"`
	TeamName    string  `json:"team_name"`
	Probability float64 `json:"probability"`
}
