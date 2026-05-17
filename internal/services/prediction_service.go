package services

import (
	"insider-league-case/internal/models"
	"sort"
)

type PredictionService interface {
	CalculatePredictions(
		teams []models.Team,
		matches []models.Match,
	) []models.Prediction
}

type MonteCarloPredictionService struct {
	matchSimulator  MatchSimulator
	standingService StandingService
}

func NewPredictionService(
	matchSimulator MatchSimulator,
	standingService StandingService,
) PredictionService {
	return &MonteCarloPredictionService{
		matchSimulator:  matchSimulator,
		standingService: standingService,
	}
}

func (s *MonteCarloPredictionService) CalculatePredictions(
	teams []models.Team,
	matches []models.Match,
) []models.Prediction {
	totalSimulations := 500

	winCounts := make(map[uint]int)

	for sim := 0; sim < totalSimulations; sim++ {
		clonedMatches := make([]models.Match, len(matches))
		copy(clonedMatches, matches)

		for i := range clonedMatches {
			if clonedMatches[i].Played {
				continue
			}

			homeTeam := findTeamByID(teams, clonedMatches[i].HomeTeamID)
			awayTeam := findTeamByID(teams, clonedMatches[i].AwayTeamID)

			clonedMatches[i] = s.matchSimulator.Simulate(
				clonedMatches[i],
				homeTeam,
				awayTeam,
			)
		}

		standings := s.standingService.CalculateStandings(
			teams,
			clonedMatches,
		)

		if len(standings) > 0 {
			champion := standings[0]
			winCounts[champion.TeamID]++
		}
	}

	var predictions []models.Prediction

	for _, team := range teams {
		probability := float64(winCounts[team.ID]) /
			float64(totalSimulations) * 100

		predictions = append(predictions, models.Prediction{
			TeamID:      team.ID,
			TeamName:    team.Name,
			Probability: probability,
		})
	}

	sort.Slice(predictions, func(i, j int) bool {
		return predictions[i].Probability > predictions[j].Probability
	})

	return predictions
}

func findTeamByID(teams []models.Team, id uint) models.Team {
	for _, team := range teams {
		if team.ID == id {
			return team
		}
	}

	return models.Team{}
}
