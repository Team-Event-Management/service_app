package models

import (
	"time"

	"github.com/google/uuid"
)

type Instance struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(255);index;not null" json:"name"`
	Lat         float64   `gorm:"type:numeric(10,7)" json:"lat"`
	Lng         float64   `gorm:"type:numeric(10,7)" json:"lng"`
	FullAddress string    `gorm:"type:text" json:"full_address"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}