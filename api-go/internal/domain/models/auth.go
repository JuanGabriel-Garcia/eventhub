package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthProps struct {
	ID        *string
	Email     *string
	Password  *string
	CreatedAt *time.Time
}

type auth struct {
	id        string
	email     string
	password  string
	createdAt time.Time
}

type Auth interface {
	GetID() string
	GetEmail() string
	GetPassword() string
	GetCreatedAt() time.Time
}

func NewAuth(props AuthProps) Auth {
	id := uuid.New().String()
	if props.ID != nil {
		id = *props.ID
	}
	createdAt := time.Now()
	if props.CreatedAt != nil {
		createdAt = *props.CreatedAt
	}
	return &auth{
		id:        id,
		email:     derefString(props.Email),
		password:  derefString(props.Password),
		createdAt: createdAt,
	}
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (a *auth) GetID() string        { return a.id }
func (a *auth) GetEmail() string     { return a.email }
func (a *auth) GetPassword() string  { return a.password }
func (a *auth) GetCreatedAt() time.Time { return a.createdAt }
