package entities

import "time"

type User struct {
	ID        string    `gorm:"primaryKey"`
	Name      string    `gorm:"not null;type:varchar(255)"`
	Email     string    `gorm:"not null;unique;type:varchar(255)"`
	Password  string    `gorm:"not null;type:varchar(255)"`
	UserType  string    `gorm:"not null;type:varchar(50);default:'participant'"`
	CreatedAt time.Time `gorm:"autoCreateTime;not null"`
}
