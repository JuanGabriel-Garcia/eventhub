package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type RegisterToEventUseCase struct {
	userRepo  repositories.UserRepository
	eventRepo repositories.IEventRepository
}

func NewRegisterToEventUseCase(userRepo repositories.UserRepository, eventRepo repositories.IEventRepository) *RegisterToEventUseCase {
	return &RegisterToEventUseCase{
		userRepo:  userRepo,
		eventRepo: eventRepo,
	}
}

type RegisterToEventUseCaseProps struct {
	UserId string
	EventId string
}

func (uc *RegisterToEventUseCase) Execute(input RegisterToEventUseCaseProps) ([]string, error) {
	event, err := uc.eventRepo.FindByID(input.EventId)
	if err != nil {
		return nil, err
	}
	user, err := uc.userRepo.FindById(input.UserId)
	if err != nil {
		return nil, err
	}

	if err := event.AddAttendee(user.GetID()); err != nil {
		return nil, err
	}

	if err := uc.eventRepo.Save(event); err != nil {
		return nil, err
	}
	
	return event.Attendees(), nil
}
