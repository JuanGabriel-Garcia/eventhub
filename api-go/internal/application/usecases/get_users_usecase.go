package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getUsersUseCase struct {
	repo repositories.UserRepository
}

func NewGetUsersUseCase(repo repositories.UserRepository) *getUsersUseCase {
	return &getUsersUseCase{repo: repo}
}

func (uc *getUsersUseCase) Execute() ([]dtos.UserResponseDTO, error) {
	users, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var dtosUsers []dtos.UserResponseDTO
	for _, user := range users {
		dtosUsers = append(dtosUsers, dtos.UserResponseDTO{
			ID:        user.GetID(),
			Name:      user.GetName(),
			Email:     user.GetEmail(),
			CreatedAt: user.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return dtosUsers, nil
}
