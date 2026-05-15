package repositories

import (
	"insider-league-case/internal/models"

	"gorm.io/gorm"
)

type MatchRepository interface {
	Create(match *models.Match) error
	FindAll() ([]models.Match, error)
	DeleteAll() error
}

type GormMatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &GormMatchRepository{db: db}
}

func (r *GormMatchRepository) Create(match *models.Match) error {
	return r.db.Create(match).Error
}

func (r *GormMatchRepository) FindAll() ([]models.Match, error) {
	var matches []models.Match
	err := r.db.Find(&matches).Error
	return matches, err
}

func (r *GormMatchRepository) DeleteAll() error {
	return r.db.Exec("DELETE FROM matches").Error
}
