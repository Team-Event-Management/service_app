package models

import "github.com/google/uuid"

type ImageEvent struct {
	IDEvent uuid.UUID `gorm:"column:id_event;type:uuid;primaryKey" json:"id_event"`
	IDImage uuid.UUID `gorm:"column:id_image;type:uuid;primaryKey" json:"id_image"`
}
