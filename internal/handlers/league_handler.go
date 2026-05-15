package handlers

import (
	"insider-league-case/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LeagueHandler struct {
	leagueService services.LeagueService
}

func NewLeagueHandler(leagueService services.LeagueService) *LeagueHandler {
	return &LeagueHandler{
		leagueService: leagueService,
	}
}

func (h *LeagueHandler) InitializeLeague(c *gin.Context) {
	if err := h.leagueService.InitializeLeague(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "League initialized successfully",
	})
}

func (h *LeagueHandler) GetStandings(c *gin.Context) {
	standings, err := h.leagueService.GetStandings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, standings)
}
