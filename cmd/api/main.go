package main

import (
	"insider-league-case/internal/database"
	"insider-league-case/internal/handlers"
	"insider-league-case/internal/repositories"
	"insider-league-case/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	teamRepo := repositories.NewTeamRepository(db)
	matchRepo := repositories.NewMatchRepository(db)

	fixtureService := services.NewFixtureService()
	standingService := services.NewStandingService()
	matchSimulator := services.NewMatchSimulator()

	leagueService := services.NewLeagueService(
		teamRepo,
		matchRepo,
		fixtureService,
		standingService,
		matchSimulator,
	)

	leagueHandler := handlers.NewLeagueHandler(leagueService)

	router := gin.Default()

	router.POST("/league/init", leagueHandler.InitializeLeague)

	router.GET("/league/table", leagueHandler.GetStandings)

	router.POST("/league/play/week/:week", leagueHandler.PlayWeek)

	router.POST("/league/play/all", leagueHandler.PlayAll)

	router.GET("/matches", leagueHandler.GetAllMatches)

	router.Run(":8080")
}
