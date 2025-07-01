package database

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/mappers"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
	mapper mappers.UserMapper
}

func NewUserRepository(db *gorm.DB, mp mappers.UserMapper) repositories.UserRepository {
	return &userRepositoryImpl{db: db, mapper: mp}
}

func (r *userRepositoryImpl) Create(user models.User) error {
	entity := r.mapper.DomainToModel(user)
	return r.db.Create(entity).Error
}

func (r *userRepositoryImpl) FindAll() ([]models.User, error) {
	var entities []entities.User
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	var users []models.User
	for _, entity := range entities {
		users = append(users, r.mapper.ModelToDomain(&entity))
	}
	return users, nil
}

func (r *userRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var entity entities.User
	if err := r.db.Where("email = ?", email).First(&entity).Error; err != nil {
		return models.NewUser(models.UserProps{}), err
	}
	user := r.mapper.ModelToDomain(&entity)
	return user, nil
}

func (r *userRepositoryImpl) FindById(id string) (models.User, error) {
	var entity entities.User
	if err := r.db.Where("id = ?", id).First(&entity).Error; err != nil {
		return models.NewUser(models.UserProps{}), err
	}
	user := r.mapper.ModelToDomain(&entity)
	return user, nil
}
