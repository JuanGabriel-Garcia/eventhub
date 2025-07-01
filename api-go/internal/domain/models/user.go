package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProps struct {
	ID        *string
	Name      *string
	Email     *string
	Password  *string
	UserType  *string
	CreatedAt *time.Time
}

type user struct {
	id        string
	name      string
	email     string
	password  string
	userType  string
	createdAt time.Time
}

type User interface {
	GetID() string
	GetName() string
	GetEmail() string
	GetPassword() string
	GetUserType() string
	GetCreatedAt() time.Time
}

func NewUser(props UserProps) User {
	id := uuid.New().String()
	if props.ID != nil {
		id = *props.ID
	}
	createdAt := time.Now()
	if props.CreatedAt != nil {
		createdAt = *props.CreatedAt
	}
	return user{
		id:        id,
		name:      derefString(props.Name),
		email:     derefString(props.Email),
		password:  derefString(props.Password),
		userType:  derefString(props.UserType),
		createdAt: createdAt,
	}
}

func (u user) GetID() string        { return u.id }
func (u user) GetName() string      { return u.name }
func (u user) GetEmail() string     { return u.email }
func (u user) GetPassword() string  { return u.password }
func (u user) GetUserType() string  { return u.userType }
func (u user) GetCreatedAt() time.Time { return u.createdAt }
