package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getEventsByCategoryUseCase struct {
    eventRepo repositories.IEventRepository
}

func NewGetEventsByCategoryUseCase(eventRepo repositories.IEventRepository) *getEventsByCategoryUseCase {
    return &getEventsByCategoryUseCase{
        eventRepo: eventRepo,
    }
}

func (uc *getEventsByCategoryUseCase) Execute(category string) ([]dtos.EventDto, error) {
    events, err := uc.eventRepo.FindByCategory(category)
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