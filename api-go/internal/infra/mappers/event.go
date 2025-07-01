package mappers

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
)

type EventMapper struct{}

func (m EventMapper) DomainToModel(event models.Event) entities.Event {
	return entities.Event{
		ID:          event.ID(),
		Name:        event.Name(),
		Location:    event.Location(),
		Date:        event.Date(),
		Description: event.Description(),
		OrganizerID: event.OrganizerID(),
		Attendees:   event.Attendees(),
		CreatedAt:   event.CreatedAt(),
		Category:    event.Category(),
		Limit: 	 event.Limit(),
	}
}

func (m EventMapper) ModelToDomain(event entities.Event) (models.Event, error) {
	domainEvent, err := models.LoadEvent(models.EventProps{
		ID:          &event.ID,
		Name:        &event.Name,
		Location:    &event.Location,
		Date:        &event.Date,
		Description: &event.Description,
		OrganizerID: &event.OrganizerID,
		Attendees:   event.Attendees,
		CreatedAt:   &event.CreatedAt,
		Category:    &event.Category,
		Limit:       &event.Limit,
	})
	if err != nil {
		return nil, err
	}
	
	return domainEvent, nil
}