package models

import (
	"time"

	"github.com/google/uuid"
)

type Stand struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	StandName       string    `gorm:"type:varchar(255);not null" json:"stand_name"`
	Lat             float64   `gorm:"type:numeric(10,7)" json:"lat"`
	Lng             float64   `gorm:"type:numeric(10,7)" json:"lng"`
	Address         string    `gorm:"type:text" json:"address"`
	StandCategoryID uuid.UUID `gorm:"type:uuid" json:"stand_category_id"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
