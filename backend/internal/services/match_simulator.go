package services

import (
	"insider-league-case/internal/models"
	"math/rand"
)

type MatchSimulator interface {
	Simulate(match models.Match, homeTeam models.Team, awayTeam models.Team) models.Match
}

type StrengthBasedMatchSimulator struct{}

func NewMatchSimulator() MatchSimulator {
	return &StrengthBasedMatchSimulator{}
}

func (s *StrengthBasedMatchSimulator) Simulate(
	match models.Match,
	homeTeam models.Team,
	awayTeam models.Team,
) models.Match {
	homeBase := homeTeam.Strength / 30
	awayBase := awayTeam.Strength / 35

	homeGoals := rand.Intn(homeBase + 2)
	awayGoals := rand.Intn(awayBase + 2)

	match.HomeGoals = homeGoals
	match.AwayGoals = awayGoals
	match.Played = true

	return match
}
