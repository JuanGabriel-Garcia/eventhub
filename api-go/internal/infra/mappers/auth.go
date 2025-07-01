package mappers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
)

type AuthMapper struct{}

func (m AuthMapper) DomainToModel(auth models.Auth) *entities.Auth {
	return &entities.Auth{
		ID:        auth.GetID(),
		Email:     auth.GetEmail(),
		Password:  auth.GetPassword(),
		CreatedAt: auth.GetCreatedAt(),
	}
}

func (m AuthMapper) ModelToDomain(entity *entities.Auth) models.Auth {
	return models.NewAuth(models.AuthProps{
		ID:        &entity.ID,
		Email:     &entity.Email,
		Password:  &entity.Password,
		CreatedAt: &entity.CreatedAt,
	})
}
