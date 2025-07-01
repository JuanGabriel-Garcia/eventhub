package usecases

import (
	"fmt"

	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getEventsUseCase struct{
	eventRepository repositories.IEventRepository
}

func NewGetEventsUseCase(eventRepository repositories.IEventRepository) *getEventsUseCase {
	return &getEventsUseCase{
		eventRepository: eventRepository,
	}
}

func (uc *getEventsUseCase) Execute() ([]dtos.EventDto, error) {
	events, err := uc.eventRepository.FindAll()
	if err != nil {
		fmt.Println("GetEventsUseCase: Retrieved events:", err)
		return nil, err
	}

	var eventDtos []dtos.EventDto
	for _, event := range events {
		eventDto := dtos.EventDto{
			ID:          event.ID(),
			Name:        event.Name(),
			Description: event.Description(),
			Location:   event.Location(),
			Date: 	  event.Date(),
			CreatedAt:  event.CreatedAt(),
			OrganizerID: event.OrganizerID(),
			Attendees: event.Attendees(),
			Category: 	 event.Category(),
			Limit:       event.Limit(),
		}
		eventDtos = append(eventDtos, eventDto)
	}
	
	return eventDtos, nil
}
