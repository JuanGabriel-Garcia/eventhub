package entities

import "time"

type Auth struct {
	ID        string    `gorm:"primaryKey"`
	Email     string
	Password  string
	CreatedAt time.Time
}
