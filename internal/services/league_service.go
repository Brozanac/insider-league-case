package services

import (
	"insider-league-case/internal/models"
	"insider-league-case/internal/repositories"
)

type LeagueService interface {
	InitializeLeague() error
	GetStandings() ([]models.Standing, error)
}

type DefaultLeagueService struct {
	teamRepo        repositories.TeamRepository
	matchRepo       repositories.MatchRepository
	fixtureService  FixtureService
	standingService StandingService
}

func NewLeagueService(
	teamRepo repositories.TeamRepository,
	matchRepo repositories.MatchRepository,
	fixtureService FixtureService,
	standingService StandingService,
) LeagueService {
	return &DefaultLeagueService{
		teamRepo:        teamRepo,
		matchRepo:       matchRepo,
		fixtureService:  fixtureService,
		standingService: standingService,
	}
}

func (s *DefaultLeagueService) InitializeLeague() error {
	s.matchRepo.DeleteAll()
	s.teamRepo.DeleteAll()

	teams := []models.Team{
		{Name: "Chelsea", Strength: 90},
		{Name: "Arsenal", Strength: 85},
		{Name: "Manchester City", Strength: 95},
		{Name: "Liverpool", Strength: 88},
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
