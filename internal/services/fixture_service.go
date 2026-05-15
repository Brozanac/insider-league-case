package services

import "insider-league-case/internal/models"

type FixtureService interface {
	GenerateFixtures(teams []models.Team) []models.Match
}

type RoundRobinFixtureService struct{}

func NewFixtureService() FixtureService {
	return &RoundRobinFixtureService{}
}

func (s *RoundRobinFixtureService) GenerateFixtures(teams []models.Team) []models.Match {
	var matches []models.Match

	week := 1

	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			matches = append(matches, models.Match{
				Week:       week,
				HomeTeamID: teams[i].ID,
				AwayTeamID: teams[j].ID,
				Played:     false,
			})

			week++
		}
	}

	for i := 0; i < len(teams); i++ {
		for j := i + 1; j < len(teams); j++ {
			matches = append(matches, models.Match{
				Week:       week,
				HomeTeamID: teams[j].ID,
				AwayTeamID: teams[i].ID,
				Played:     false,
			})

			week++
		}
	}

	return matches
}
