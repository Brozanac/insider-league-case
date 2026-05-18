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
	if len(teams) != 4 {
		return []models.Match{}
	}

	return []models.Match{
		// First half
		{
			Week:       1,
			HomeTeamID: teams[0].ID,
			AwayTeamID: teams[1].ID,
			Played:     false,
		},
		{
			Week:       1,
			HomeTeamID: teams[2].ID,
			AwayTeamID: teams[3].ID,
			Played:     false,
		},

		{
			Week:       2,
			HomeTeamID: teams[0].ID,
			AwayTeamID: teams[2].ID,
			Played:     false,
		},
		{
			Week:       2,
			HomeTeamID: teams[1].ID,
			AwayTeamID: teams[3].ID,
			Played:     false,
		},

		{
			Week:       3,
			HomeTeamID: teams[0].ID,
			AwayTeamID: teams[3].ID,
			Played:     false,
		},
		{
			Week:       3,
			HomeTeamID: teams[1].ID,
			AwayTeamID: teams[2].ID,
			Played:     false,
		},

		// Second half — reverse home/away
		{
			Week:       4,
			HomeTeamID: teams[1].ID,
			AwayTeamID: teams[0].ID,
			Played:     false,
		},
		{
			Week:       4,
			HomeTeamID: teams[3].ID,
			AwayTeamID: teams[2].ID,
			Played:     false,
		},

		{
			Week:       5,
			HomeTeamID: teams[2].ID,
			AwayTeamID: teams[0].ID,
			Played:     false,
		},
		{
			Week:       5,
			HomeTeamID: teams[3].ID,
			AwayTeamID: teams[1].ID,
			Played:     false,
		},

		{
			Week:       6,
			HomeTeamID: teams[3].ID,
			AwayTeamID: teams[0].ID,
			Played:     false,
		},
		{
			Week:       6,
			HomeTeamID: teams[2].ID,
			AwayTeamID: teams[1].ID,
			Played:     false,
		},
	}
}
