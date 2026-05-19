package handlers

import (
	"insider-league-case/internal/models"
	"insider-league-case/internal/services"
	"net/http"
	"strconv"

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

func (h *LeagueHandler) PlayWeek(c *gin.Context) {
	weekParam := c.Param("week")

	week, err := strconv.Atoi(weekParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid week parameter",
		})
		return
	}

	if err := h.leagueService.PlayWeek(week); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Week played successfully",
		"week":    week,
	})
}

func (h *LeagueHandler) PlayAll(c *gin.Context) {
	if err := h.leagueService.PlayAll(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "All remaining matches played successfully",
	})
}

func (h *LeagueHandler) GetAllMatches(c *gin.Context) {
	matches, err := h.leagueService.GetAllMatches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, matches)
}

func (h *LeagueHandler) GetPredictions(c *gin.Context) {
	predictions, err := h.leagueService.GetPredictions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, predictions)
}

func (h *LeagueHandler) UpdateMatchResult(c *gin.Context) {
	matchIDParam := c.Param("id")

	matchID, err := strconv.Atoi(matchIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid match id",
		})
		return
	}

	var request models.UpdateMatchRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if request.HomeGoals < 0 || request.AwayGoals < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Goals cannot be negative",
		})
		return
	}

	if err := h.leagueService.UpdateMatchResult(
		uint(matchID),
		request.HomeGoals,
		request.AwayGoals,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Match result updated successfully",
	})
}
