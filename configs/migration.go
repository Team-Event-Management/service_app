package configs

import (
	"event_management/internal/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Role{},
		&models.User{},
	)
}
