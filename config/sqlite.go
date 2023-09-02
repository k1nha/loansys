package config

import (
	"github.com/lucascmpus/lend/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDBProvider() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open("./db/dev.db"), &gorm.Config{})

	if err != nil {
		logger.ErrorF("Error creating database connection: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Role{})

	if err != nil {
		logger.ErrorF("SQLite Automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
