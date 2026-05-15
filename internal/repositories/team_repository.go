package repositories

import (
	"insider-league-case/internal/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	Create(team *models.Team) error
	FindAll() ([]models.Team, error)
	DeleteAll() error
}

type GormTeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &GormTeamRepository{db: db}
}

func (r *GormTeamRepository) Create(team *models.Team) error {
	return r.db.Create(team).Error
}

func (r *GormTeamRepository) FindAll() ([]models.Team, error) {
	var teams []models.Team
	err := r.db.Find(&teams).Error
	return teams, err
}

func (r *GormTeamRepository) DeleteAll() error {
	return r.db.Exec("DELETE FROM teams").Error
}
