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
	if len(teams)%2 != 0 {
		return []models.Match{}
	}

	var matches []models.Match

	n := len(teams)
	rounds := n - 1
	matchesPerRound := n / 2

	// Work on a copy so we do not mutate the original teams slice.
	rotatingTeams := make([]models.Team, n)
	copy(rotatingTeams, teams)

	// First half of the season.
	for round := 0; round < rounds; round++ {
		week := round + 1

		for i := 0; i < matchesPerRound; i++ {
			home := rotatingTeams[i]
			away := rotatingTeams[n-1-i]

			// Alternate home/away a little so the first team does not always stay home.
			if round%2 == 0 {
				matches = append(matches, models.Match{
					Week:       week,
					HomeTeamID: home.ID,
					AwayTeamID: away.ID,
					Played:     false,
				})
			} else {
				matches = append(matches, models.Match{
					Week:       week,
					HomeTeamID: away.ID,
					AwayTeamID: home.ID,
					Played:     false,
				})
			}
		}

		// Rotate all teams except the first one.
		last := rotatingTeams[n-1]
		copy(rotatingTeams[2:], rotatingTeams[1:n-1])
		rotatingTeams[1] = last
	}

	// Second half of the season: reverse home/away.
	firstHalfCount := len(matches)

	for i := 0; i < firstHalfCount; i++ {
		matches = append(matches, models.Match{
			Week:       matches[i].Week + rounds,
			HomeTeamID: matches[i].AwayTeamID,
			AwayTeamID: matches[i].HomeTeamID,
			Played:     false,
		})
	}

	return matches
}
