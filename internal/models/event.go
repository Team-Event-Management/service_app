package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID        	uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	NameEvent 	string    `gorm:"type:varchar(255)" json:"name_event"`
	Description string    `gorm:"type:text" json:"description"`
	Status      int       `gorm:"type:int;" json:"status"`
	Location  	string    `gorm:"type:text" json:"location"`
	EventImages  []Image   `gorm:"many2many:image_events;joinForeignKey:ID;joinReferences:ID;joinForeignKey:id_event;joinReferences:id_image" json:"event_images"`
	CreatedAt 	time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt 	time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}