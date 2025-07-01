package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getAuthsUseCase struct {
	repo repositories.AuthRepository
}

func NewGetAuthsUseCase(repo repositories.AuthRepository) *getAuthsUseCase {
	return &getAuthsUseCase{repo: repo}
}

func (uc *getAuthsUseCase) Execute() ([]dtos.AuthResponseDTO, error) {
	auths, err := uc.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var dtosAuths []dtos.AuthResponseDTO
	for _, auth := range auths {
		dtosAuths = append(dtosAuths, dtos.AuthResponseDTO{
			ID:        auth.GetID(),
			Email:     auth.GetEmail(),
			CreatedAt: auth.GetCreatedAt().Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return dtosAuths, nil
}
