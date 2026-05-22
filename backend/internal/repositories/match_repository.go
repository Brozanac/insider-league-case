package repositories

import (
	"insider-league-case/internal/models"

	"gorm.io/gorm"
)

type MatchRepository interface {
	Create(match *models.Match) error
	FindAll() ([]models.Match, error)
	FindByID(id uint) (models.Match, error)
	FindByWeek(week int) ([]models.Match, error)
	FindUnplayed() ([]models.Match, error)
	Update(match *models.Match) error
	DeleteAll() error
	ResetAutoIncrement() error
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
	err := r.db.
		Order("week ASC").
		Order("id ASC").
		Find(&matches).Error

	return matches, err
}

func (r *GormMatchRepository) FindByWeek(week int) ([]models.Match, error) {
	var matches []models.Match
	err := r.db.
		Where("week = ?", week).
		Order("id ASC").
		Find(&matches).Error

	return matches, err
}

func (r *GormMatchRepository) FindUnplayed() ([]models.Match, error) {
	var matches []models.Match
	err := r.db.
		Where("played = ?", false).
		Order("week ASC").
		Order("id ASC").
		Find(&matches).Error

	return matches, err
}

func (r *GormMatchRepository) Update(match *models.Match) error {
	return r.db.Save(match).Error
}

func (r *GormMatchRepository) DeleteAll() error {
	return r.db.Exec("DELETE FROM matches").Error
}

func (r *GormMatchRepository) ResetAutoIncrement() error {
	return r.db.Exec("DELETE FROM sqlite_sequence WHERE name='matches'").Error
}

func (r *GormMatchRepository) FindByID(id uint) (models.Match, error) {
	var match models.Match
	err := r.db.First(&match, id).Error
	return match, err
}
