package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeDBProvider()

	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}

	return err
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
