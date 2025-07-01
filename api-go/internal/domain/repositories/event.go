package repositories

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
)

type IEventRepository interface {
	FindByID(id string) (models.Event, error)
	FindAll() ([]models.Event, error)
	FindByAttendee(userID string) ([]models.Event, error)
	FindByOrganizerID(organizerID string) ([]models.Event, error)
	FindEventByOrganizerID(eventID, organizerID string) (models.Event, error)
	FindByCategory(category string) ([]models.Event, error)
	FindByTerm(term string) ([]models.Event, error)
	Save(event models.Event) error
	Delete(id string) error
}