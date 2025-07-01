package usecases

import (
	"errors"
	"log"
	
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/services"
	"github.com/Gabriel-Schiestl/api-go/internal/utils"
)

type loginUseCase struct {
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
	jwtService services.IJWTService
}

func NewLoginUseCase(authRepo repositories.AuthRepository, userRepo repositories.UserRepository, jwtService services.IJWTService) *loginUseCase {
	return &loginUseCase{authRepo: authRepo, userRepo: userRepo, jwtService: jwtService}
}

func (uc *loginUseCase) Execute(props dtos.LoginDto) (*string, error) {
	log.Printf("LoginUseCase - Attempting login for email: %s", props.Email)
	
	user, err := uc.userRepo.FindByEmail(props.Email)
	if err != nil {
		log.Printf("LoginUseCase - User not found for email %s: %v", props.Email, err)
		return nil, err
	}

	log.Printf("LoginUseCase - Found user: ID=%s, Name=%s, Email=%s", user.GetID(), user.GetName(), user.GetEmail())

	// Verificar a senha que está na tabela users
	if !utils.CheckPasswordHash(props.Password, user.GetPassword()) {
		log.Printf("LoginUseCase - Invalid password for user %s", user.GetEmail())
		return nil, errors.New("credenciais inválidas")
	}

	log.Printf("LoginUseCase - Password verified, generating token for user ID: %s", user.GetID())
	token, err := uc.jwtService.GenerateToken(user.GetID())
	if err != nil {
		log.Printf("LoginUseCase - Error generating token: %v", err)
		return nil, err
	}

	log.Printf("LoginUseCase - Token generated successfully for user %s", user.GetID())
	return token, nil
}
