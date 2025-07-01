package usecases

import (
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type CancelEventSubscriptionUseCase struct {
	userRepo  repositories.UserRepository
	eventRepo repositories.IEventRepository
}

func NewCancelEventSubscriptionUseCase(userRepo repositories.UserRepository, eventRepo repositories.IEventRepository) *CancelEventSubscriptionUseCase {
	return &CancelEventSubscriptionUseCase{
		userRepo:  userRepo,
		eventRepo: eventRepo,
	}
}

type CancelEventSubscriptionUseCaseProps struct {
	UserId  string
	EventId string
}

func (uc *CancelEventSubscriptionUseCase) Execute(input CancelEventSubscriptionUseCaseProps) ([]string, error) {
	event, err := uc.eventRepo.FindByID(input.EventId)
	if err != nil {
		return nil, err
	}
	
	user, err := uc.userRepo.FindById(input.UserId)
	if err != nil {
		return nil, err
	}

	if err := event.CancelSubscription(user.GetID()); err != nil {
		return nil, err
	}

	if err := uc.eventRepo.Save(event); err != nil {
		return nil, err
	}
	
	return event.Attendees(), nil
}