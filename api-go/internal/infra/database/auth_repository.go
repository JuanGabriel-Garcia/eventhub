package database

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/mappers"
	"gorm.io/gorm"
)

type authRepositoryImpl struct {
	db *gorm.DB
	mapper mappers.AuthMapper
}

func NewAuthRepository(db *gorm.DB, mapper mappers.AuthMapper) repositories.AuthRepository {
	return &authRepositoryImpl{db: db, mapper: mapper}
}

func (r *authRepositoryImpl) Create(auth models.Auth) error {
	entity := r.mapper.DomainToModel(auth)
	return r.db.Create(entity).Error
}

func (r *authRepositoryImpl) FindAll() ([]models.Auth, error) {
	var entities []entities.Auth
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	var auths []models.Auth
	for _, entity := range entities {
		auths = append(auths, r.mapper.ModelToDomain(&entity))
	}
	return auths, nil
}

func (r *authRepositoryImpl) FindByEmail(email string) (models.Auth, error) {
	var entity entities.Auth
	if err := r.db.Where("email = ?", email).First(&entity).Error; err != nil {
		return models.NewAuth(models.AuthProps{}), err
	}
	auth := r.mapper.ModelToDomain(&entity)
	return auth, nil
}
