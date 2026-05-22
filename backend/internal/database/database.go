package database

import (
	"insider-league-case/internal/models"
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Open("data/league.db"), &gorm.Config{})
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
