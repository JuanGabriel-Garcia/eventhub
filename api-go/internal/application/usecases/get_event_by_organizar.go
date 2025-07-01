package usecases

import (
	"sync"

	"github.com/Gabriel-Schiestl/api-go/internal/application/dtos"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
)

type getEventByOrganizerUseCase struct {
	eventRepo repositories.IEventRepository
	userRepo repositories.UserRepository
}

func NewGetEventByOrganizerUseCase(eventRepo repositories.IEventRepository, userRepo repositories.UserRepository) *getEventByOrganizerUseCase {
	return &getEventByOrganizerUseCase{
		eventRepo: eventRepo,
		userRepo: userRepo,
	}
}

type GetEventByOrganizerUseCaseProps struct {
	OrganizerId string
	EventId     string
}

func (uc *getEventByOrganizerUseCase) Execute(props GetEventByOrganizerUseCaseProps) (dtos.EventWithAttendeesDto, error) {
	event, err := uc.eventRepo.FindEventByOrganizerID(props.EventId, props.OrganizerId)
	if err != nil {
		return dtos.EventWithAttendeesDto{}, err
	}

	usersChan := make(chan models.User)
	wg := sync.WaitGroup{}

	for _, attendeeId := range event.Attendees() {
		wg.Add(1)
		go func(attendeeId string) {
			defer wg.Done()

			user, err := uc.userRepo.FindById(attendeeId)
			if err != nil {
				usersChan <- nil
				return
			}

			usersChan <- user
		}(attendeeId)
	}

	go func() {
		wg.Wait()
		close(usersChan)
	}()

	var users []dtos.UserResponseDTO
	for user := range usersChan {
		if user != nil {
			users = append(users, dtos.UserResponseDTO{
				ID:        user.GetID(),
				Email:     user.GetEmail(),
				Name: user.GetName(),
				CreatedAt: user.GetCreatedAt().String(),
			})
		}
	}

	eventDto := dtos.EventWithAttendeesDto{
		ID: 		event.ID(),
		Name: 		event.Name(),
		Description: event.Description(),
		Location: 	event.Location(),
		Date: 		event.Date(),
		OrganizerID: event.OrganizerID(),
		Attendees: 	users,
		CreatedAt: event.CreatedAt(),
		Category: 	 event.Category(),
		Limit:       event.Limit(),
	}
	
	return eventDto, nil
}