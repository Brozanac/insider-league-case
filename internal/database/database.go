package database

import (
	"insider-league-case/internal/models"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("league.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Team{},
		&models.Match{},
	)

	if err != nil {
		return nil, err
	}

	return db, nil
}
