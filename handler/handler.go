package handler

import (
	"github.com/lucascmpus/lend/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() (*config.Logger, *gorm.DB) {
	logger = config.GetLogger("handler")
	db = config.GetDB()

	return logger, db
}
