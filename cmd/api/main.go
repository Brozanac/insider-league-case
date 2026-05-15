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

	leagueService := services.NewLeagueService(
		teamRepo,
		matchRepo,
		fixtureService,
		standingService,
	)

	leagueHandler := handlers.NewLeagueHandler(leagueService)

	router := gin.Default()

	router.POST("/league/init", leagueHandler.InitializeLeague)
	router.GET("/league/table", leagueHandler.GetStandings)

	router.Run(":8080")
}
