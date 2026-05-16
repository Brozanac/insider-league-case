package repositories

import (
	"insider-league-case/internal/models"

	"gorm.io/gorm"
)

type TeamRepository interface {
	Create(team *models.Team) error
	FindByID(id uint) (models.Team, error)
	FindAll() ([]models.Team, error)
	DeleteAll() error
	ResetAutoIncrement() error
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

func (r *GormTeamRepository) FindByID(id uint) (models.Team, error) {
	var team models.Team
	err := r.db.First(&team, id).Error
	return team, err
}

func (r *GormTeamRepository) ResetAutoIncrement() error {
	return r.db.Exec("DELETE FROM sqlite_sequence WHERE name='teams'").Error
}
