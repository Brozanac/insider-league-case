package services

import (
	"insider-league-case/internal/models"
	"insider-league-case/internal/repositories"
)

type LeagueService interface {
	InitializeLeague() error
	GetStandings() ([]models.Standing, error)
	PlayWeek(week int) error
	PlayAll() error
	GetAllMatches() ([]models.Match, error)
	GetPredictions() ([]models.Prediction, error)
}

type DefaultLeagueService struct {
	teamRepo          repositories.TeamRepository
	matchRepo         repositories.MatchRepository
	fixtureService    FixtureService
	standingService   StandingService
	matchSimulator    MatchSimulator
	predictionService PredictionService
}

func NewLeagueService(
	teamRepo repositories.TeamRepository,
	matchRepo repositories.MatchRepository,
	fixtureService FixtureService,
	standingService StandingService,
	matchSimulator MatchSimulator,
	predictionService PredictionService,
) LeagueService {
	return &DefaultLeagueService{
		teamRepo:          teamRepo,
		matchRepo:         matchRepo,
		fixtureService:    fixtureService,
		standingService:   standingService,
		matchSimulator:    matchSimulator,
		predictionService: predictionService,
	}
}

func (s *DefaultLeagueService) InitializeLeague() error {
	s.matchRepo.DeleteAll()
	s.teamRepo.DeleteAll()
	s.matchRepo.ResetAutoIncrement()
	s.teamRepo.ResetAutoIncrement()

	teams := []models.Team{
		{Name: "Arsenal", Strength: 90},
		{Name: "Aston Villa", Strength: 82},
		{Name: "Bournemouth", Strength: 74},
		{Name: "Brentford", Strength: 73},
		{Name: "Brighton & Hove Albion", Strength: 78},
		{Name: "Burnley", Strength: 65},
		{Name: "Chelsea", Strength: 86},
		{Name: "Crystal Palace", Strength: 76},
		{Name: "Everton", Strength: 72},
		{Name: "Fulham", Strength: 74},
		{Name: "Leeds United", Strength: 68},
		{Name: "Liverpool", Strength: 91},
		{Name: "Manchester City", Strength: 94},
		{Name: "Manchester United", Strength: 84},
		{Name: "Newcastle United", Strength: 83},
		{Name: "Nottingham Forest", Strength: 75},
		{Name: "Sunderland", Strength: 66},
		{Name: "Tottenham Hotspur", Strength: 81},
		{Name: "West Ham United", Strength: 77},
		{Name: "Wolverhampton Wanderers", Strength: 70},
	}

	for i := range teams {
		if err := s.teamRepo.Create(&teams[i]); err != nil {
			return err
		}
	}

	savedTeams, err := s.teamRepo.FindAll()
	if err != nil {
		return err
	}

	matches := s.fixtureService.GenerateFixtures(savedTeams)

	for i := range matches {
		if err := s.matchRepo.Create(&matches[i]); err != nil {
			return err
		}
	}

	return nil
}

func (s *DefaultLeagueService) GetStandings() ([]models.Standing, error) {
	teams, err := s.teamRepo.FindAll()
	if err != nil {
		return nil, err
	}

	matches, err := s.matchRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return s.standingService.CalculateStandings(teams, matches), nil
}

func (s *DefaultLeagueService) PlayWeek(week int) error {
	matches, err := s.matchRepo.FindByWeek(week)
	if err != nil {
		return err
	}

	for i := range matches {
		if matches[i].Played {
			continue
		}

		homeTeam, err := s.teamRepo.FindByID(matches[i].HomeTeamID)
		if err != nil {
			return err
		}

		awayTeam, err := s.teamRepo.FindByID(matches[i].AwayTeamID)
		if err != nil {
			return err
		}

		simulatedMatch := s.matchSimulator.Simulate(
			matches[i],
			homeTeam,
			awayTeam,
		)

		if err := s.matchRepo.Update(&simulatedMatch); err != nil {
			return err
		}
	}

	return nil
}

func (s *DefaultLeagueService) PlayAll() error {
	matches, err := s.matchRepo.FindUnplayed()
	if err != nil {
		return err
	}

	for i := range matches {
		homeTeam, err := s.teamRepo.FindByID(matches[i].HomeTeamID)
		if err != nil {
			return err
		}

		awayTeam, err := s.teamRepo.FindByID(matches[i].AwayTeamID)
		if err != nil {
			return err
		}

		simulatedMatch := s.matchSimulator.Simulate(
			matches[i],
			homeTeam,
			awayTeam,
		)

		if err := s.matchRepo.Update(&simulatedMatch); err != nil {
			return err
		}
	}

	return nil
}

func (s *DefaultLeagueService) GetAllMatches() ([]models.Match, error) {
	return s.matchRepo.FindAll()
}

func (s *DefaultLeagueService) GetPredictions() ([]models.Prediction, error) {
	teams, err := s.teamRepo.FindAll()
	if err != nil {
		return nil, err
	}

	matches, err := s.matchRepo.FindAll()
	if err != nil {
		return nil, err
	}

	return s.predictionService.CalculatePredictions(
		teams,
		matches,
	), nil
}
