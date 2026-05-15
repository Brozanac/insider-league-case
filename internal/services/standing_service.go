package services

import (
	"insider-league-case/internal/models"
	"sort"
)

type StandingService interface {
	CalculateStandings(teams []models.Team, matches []models.Match) []models.Standing
}

type PremierLeagueStandingService struct{}

func NewStandingService() StandingService {
	return &PremierLeagueStandingService{}
}

func (s *PremierLeagueStandingService) CalculateStandings(
	teams []models.Team,
	matches []models.Match,
) []models.Standing {
	table := make(map[uint]*models.Standing)

	for _, team := range teams {
		table[team.ID] = &models.Standing{
			TeamID:   team.ID,
			TeamName: team.Name,
		}
	}

	for _, match := range matches {
		if !match.Played {
			continue
		}

		home := table[match.HomeTeamID]
		away := table[match.AwayTeamID]

		home.Played++
		away.Played++

		home.GoalsFor += match.HomeGoals
		home.GoalsAgainst += match.AwayGoals

		away.GoalsFor += match.AwayGoals
		away.GoalsAgainst += match.HomeGoals

		if match.HomeGoals > match.AwayGoals {
			home.Won++
			away.Lost++
			home.Points += 3
		} else if match.HomeGoals < match.AwayGoals {
			away.Won++
			home.Lost++
			away.Points += 3
		} else {
			home.Drawn++
			away.Drawn++
			home.Points++
			away.Points++
		}
	}

	var standings []models.Standing

	for _, standing := range table {
		standing.GoalDiff = standing.GoalsFor - standing.GoalsAgainst
		standings = append(standings, *standing)
	}

	sort.Slice(standings, func(i, j int) bool {
		if standings[i].Points != standings[j].Points {
			return standings[i].Points > standings[j].Points
		}

		if standings[i].GoalDiff != standings[j].GoalDiff {
			return standings[i].GoalDiff > standings[j].GoalDiff
		}

		return standings[i].GoalsFor > standings[j].GoalsFor
	})

	return standings
}
