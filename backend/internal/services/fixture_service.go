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
	if len(teams) < 2 {
		return []models.Match{}
	}

	players := append([]models.Team{}, teams...)

	hasBye := false
	if len(players)%2 != 0 {
		hasBye = true
		players = append(players, models.Team{})
	}

	teamCount := len(players)
	roundCount := teamCount - 1
	matchesPerRound := teamCount / 2

	var matches []models.Match

	for round := 0; round < roundCount; round++ {
		for i := 0; i < matchesPerRound; i++ {
			home := players[i]
			away := players[teamCount-1-i]

			if hasBye && (home.ID == 0 || away.ID == 0) {
				continue
			}

			matches = append(matches, models.Match{
				Week:       round + 1,
				HomeTeamID: home.ID,
				AwayTeamID: away.ID,
				Played:     false,
			})
		}

		players = rotateTeams(players)
	}

	firstHalfMatchCount := len(matches)

	for i := 0; i < firstHalfMatchCount; i++ {
		firstLeg := matches[i]

		matches = append(matches, models.Match{
			Week:       firstLeg.Week + roundCount,
			HomeTeamID: firstLeg.AwayTeamID,
			AwayTeamID: firstLeg.HomeTeamID,
			Played:     false,
		})
	}

	return matches
}

func rotateTeams(teams []models.Team) []models.Team {
	n := len(teams)

	rotated := make([]models.Team, n)

	rotated[0] = teams[0]
	rotated[1] = teams[n-1]

	for i := 2; i < n; i++ {
		rotated[i] = teams[i-1]
	}

	return rotated
}
