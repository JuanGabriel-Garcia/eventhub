package entities

import (
	"time"

	"github.com/Gabriel-Schiestl/api-go/internal/utils"
)

type Event struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null;type:varchar(255)"`
	Location    string `gorm:"not null;type:varchar(255)"`
	Date        time.Time `gorm:"not null"`
	Description string `gorm:"type:text"`
	OrganizerID string `gorm:"not null;type:varchar(255)"`
	Attendees   utils.StringArray `gorm:"type:json"`
	CreatedAt   time.Time `gorm:"autoCreateTime;not null"`
	Category	string `gorm:"not null;type:varchar(255)"`
	Limit       int `gorm:"not null;default:0"`
}