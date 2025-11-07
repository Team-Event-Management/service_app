package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(255)" json:"name"`
	Email       string     `gorm:"type:varchar(255);uniqueIndex" json:"email"`
	Password    string     `gorm:"type:varchar(255)" json:"password"`
	PhoneNumber *string    `gorm:"type:varchar(20)" json:"phone_number"`
	DateOfBirth *time.Time `gorm:"type:date;" json:"date_of_birth"`
	Age         *int       `gorm:"type:int;" json:"age"`
	Status      int        `gorm:"type:int;" json:"status"`
	RoleID      uuid.UUID  `gorm:"type:uuid"`
	Role        Role       `gorm:"foreignKey:RoleID"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
