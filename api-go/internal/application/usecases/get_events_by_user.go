package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getEventsByUserUseCase struct {
	userRepo  repositories.UserRepository
	eventRepo repositories.IEventRepository
}

func NewGetEventsByUserUseCase(userRepo repositories.UserRepository, eventRepo repositories.IEventRepository) *getEventsByUserUseCase {
	return &getEventsByUserUseCase{
		userRepo:  userRepo,
		eventRepo: eventRepo,
	}
}

func (uc *getEventsByUserUseCase) Execute(userId string) ([]dtos.EventDto, error) {
	user, err := uc.userRepo.FindById(userId)
	if err != nil {
		return nil, err
	}

	events, err := uc.eventRepo.FindByAttendee(user.GetID())
	if err != nil {
		return nil, err
	}

	var eventDtos []dtos.EventDto
	for _, event := range events {
		eventDtos = append(eventDtos, dtos.EventDto{
			ID:          event.ID(),
			Name: 	event.Name(),
			Description: event.Description(),
			Location:    event.Location(),
			Date: event.Date(),
			OrganizerID: event.OrganizerID(),
			Attendees: event.Attendees(),
			CreatedAt:  event.CreatedAt(),
			Category: 	 event.Category(),
			Limit: 	 event.Limit(),
		})
	}
	return eventDtos, nil
}
