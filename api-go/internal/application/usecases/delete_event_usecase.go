package usecases

import (
	"fmt"

	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/go-clarch/domain/exceptions"
)

type deleteEventUseCase struct {
	eventRepository repositories.IEventRepository
}

func NewDeleteEventUseCase(eventRepository repositories.IEventRepository) *deleteEventUseCase {
	return &deleteEventUseCase{
		eventRepository: eventRepository,
	}
}

type DeleteEventProps struct {
	EventID     string
	OrganizerID string
}

func (uc *deleteEventUseCase) Execute(props DeleteEventProps) (struct{}, error) {
	event, err := uc.eventRepository.FindByID(props.EventID)
	if err != nil {
		return struct{}{}, err
	}

	if event.OrganizerID() != props.OrganizerID {
		return struct{}{}, exceptions.NewBusinessException(fmt.Sprintf("User %s is not authorized to delete event %s", props.OrganizerID, props.EventID))
	}

	deleteErr := uc.eventRepository.Delete(props.EventID)
	if deleteErr != nil {
		return struct{}{}, deleteErr
	}

	return struct{}{}, nil
}
