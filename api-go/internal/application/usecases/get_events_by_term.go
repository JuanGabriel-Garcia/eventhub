package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getEventsByTermUseCase struct {
    eventRepo repositories.IEventRepository
}

func NewGetEventsByTermUseCase(eventRepo repositories.IEventRepository) *getEventsByTermUseCase {
    return &getEventsByTermUseCase{
        eventRepo: eventRepo,
    }
}

func (uc *getEventsByTermUseCase) Execute(term string) ([]dtos.EventDto, error) {
    events, err := uc.eventRepo.FindByTerm(term)
    if err != nil {
        return nil, err
    }

    var eventDtos []dtos.EventDto
    for _, event := range events {
        eventDtos = append(eventDtos, dtos.EventDto{
            ID:          event.ID(),
            Name:        event.Name(),
            Description: event.Description(),
            Location:    event.Location(),
            Date:        event.Date(),
            OrganizerID: event.OrganizerID(),
            Attendees:   event.Attendees(),
            CreatedAt:   event.CreatedAt(),
            Category:    event.Category(),
            Limit:       event.Limit(),
        })
    }

    return eventDtos, nil
}