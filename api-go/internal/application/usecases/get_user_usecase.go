package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getUserUseCase struct {
	repo repositories.UserRepository
}

func NewGetUserUseCase(repo repositories.UserRepository) *getUserUseCase {
	return &getUserUseCase{repo: repo}
}

func (uc *getUserUseCase) Execute(id string) (dtos.UserResponseDTO, error) {
	user, err := uc.repo.FindById(id)
	if err != nil {
		return dtos.UserResponseDTO{}, err
	}

	return dtos.UserResponseDTO{
		ID:        user.GetID(),
		Name:      user.GetName(),
		Email:     user.GetEmail(),
		CreatedAt: user.GetCreatedAt().String(),
	}, nil
}
