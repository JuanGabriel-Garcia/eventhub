package database

import (
	"fmt"
	"log"

	"github.com/Gabriel-Schiestl/api-go/internal/domain/models"
	"github.com/Gabriel-Schiestl/api-go/internal/domain/repositories"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/entities"
	"github.com/Gabriel-Schiestl/api-go/internal/infra/mappers"
	"gorm.io/gorm"
)

var errorLoadingEvent = "Error loading event: %v"

type eventRepositoryImpl struct {
	db     *gorm.DB
	mapper mappers.EventMapper
}

func NewEventRepository(db *gorm.DB, mapper mappers.EventMapper) repositories.IEventRepository {
	return eventRepositoryImpl{
		db:     db,
		mapper: mapper,
	}
}

func (r eventRepositoryImpl) FindByID(id string) (models.Event, error) {
	var event entities.Event
	if err := r.db.First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("event with ID %s not found", id)
		}

		return nil, fmt.Errorf("error retrieving event with ID %s: %v", id, err)
	}

	domain, err := r.mapper.ModelToDomain(event)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r eventRepositoryImpl) FindAll() ([]models.Event, error) {
	var events []entities.Event

	if err := r.db.Find(&events).Error; err != nil {
		return nil, fmt.Errorf("error retrieving events: %v", err)
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("no events found")
	}

	var domainEvents []models.Event
	for _, event := range events {
		domain, err := r.mapper.ModelToDomain(event)
		if err != nil {
			fmt.Printf(errorLoadingEvent, err)
			return nil, err
		}

		domainEvents = append(domainEvents, domain)
	}

	return domainEvents, nil
}

func (r eventRepositoryImpl) FindByAttendee(userID string) ([]models.Event, error) {
	var events []entities.Event

	query := `
        SELECT * FROM events 
        WHERE EXISTS (
            SELECT 1 FROM json_array_elements_text(attendees) AS attendee 
            WHERE attendee = ?
        )
    `

	if err := r.db.Raw(query, userID).Scan(&events).Error; err != nil {
		return nil, fmt.Errorf("error retrieving events for user ID %s: %v", userID, err)
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("no events found for user ID %s", userID)
	}

	var domainEvents []models.Event
	for _, event := range events {
		domain, err := r.mapper.ModelToDomain(event)
		if err != nil {
			fmt.Printf(errorLoadingEvent, err)
			return nil, err
		}

		domainEvents = append(domainEvents, domain)
	}

	return domainEvents, nil
}

func (r eventRepositoryImpl) FindByOrganizerID(organizerID string) ([]models.Event, error) {
	var events []entities.Event

	log.Printf("FindByOrganizerID - Searching for events with organizer_id = %s", organizerID)

	if err := r.db.Where("organizer_id = ?", organizerID).Find(&events).Error; err != nil {
		log.Printf("FindByOrganizerID - Database error: %v", err)
		return nil, fmt.Errorf("error retrieving events for organizer ID %s: %v", organizerID, err)
	}

	log.Printf("FindByOrganizerID - Found %d events in database for organizer %s", len(events), organizerID)
	
	if len(events) == 0 {
		log.Printf("FindByOrganizerID - No events found for organizer ID %s", organizerID)
		return nil, fmt.Errorf("no events found for organizer ID %s", organizerID)
	}

	// Log details of each event found
	for i, event := range events {
		log.Printf("Event %d: ID=%s, Name=%s, OrganizerID=%s", i+1, event.ID, event.Name, event.OrganizerID)
	}

	var domainEvents []models.Event
	for _, event := range events {
		domain, err := r.mapper.ModelToDomain(event)
		if err != nil {
			fmt.Printf(errorLoadingEvent, err)
			return nil, err
		}

		domainEvents = append(domainEvents, domain)
	}

	return domainEvents, nil
}

func (r eventRepositoryImpl) FindEventByOrganizerID(eventID, organizerID string) (models.Event, error) {
	var event entities.Event

	if err := r.db.Where("id = ? AND organizer_id = ?", eventID, organizerID).First(&event).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("Event with ID %s not found for organizer ID %s", eventID, organizerID)
		}
		return nil, fmt.Errorf("Error retrieving event with ID %s for organizer ID %s: %v", eventID, organizerID, err)
	}

	domain, err := r.mapper.ModelToDomain(event)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r eventRepositoryImpl) FindByCategory(category string) ([]models.Event, error) {
	var events []entities.Event

	if err := r.db.Where("category = ?", category).Find(&events).Error; err != nil {
		return nil, fmt.Errorf("Error retrieving events for category %s: %v", category, err)
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("No events found for category %s", category)
	}

	var domainEvents []models.Event
	for _, event := range events {
		domain, err := r.mapper.ModelToDomain(event)
		if err != nil {
			fmt.Printf(errorLoadingEvent, err)
			return nil, err
		}

		domainEvents = append(domainEvents, domain)
	}

	return domainEvents, nil
}

func (r eventRepositoryImpl) FindByTerm(term string) ([]models.Event, error) {
	var events []entities.Event

	if err := r.db.Where("name LIKE ? OR description LIKE ?", "%"+term+"%", "%"+term+"%").Find(&events).Error; err != nil {
		return nil, fmt.Errorf("Error retrieving events by term %s: %v", term, err)
	}

	if len(events) == 0 {
		return nil, fmt.Errorf("No events found for term %s", term)
	}

	var domainEvents []models.Event
	for _, event := range events {
		domain, err := r.mapper.ModelToDomain(event)
		if err != nil {
			fmt.Printf(errorLoadingEvent, err)
			return nil, err
		}

		domainEvents = append(domainEvents, domain)
	}

	return domainEvents, nil
}

func (r eventRepositoryImpl) Save(event models.Event) error {
	entity := r.mapper.DomainToModel(event)
	if err := r.db.Save(&entity).Error; err != nil {
		return fmt.Errorf("Error saving event: %v", err)
	}

	return nil
}

func (r eventRepositoryImpl) Delete(id string) error {
	var event entities.Event
	if err := r.db.First(&event, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Event with ID %s not found", id)
		}
		return fmt.Errorf("Error retrieving event with ID %s: %v", id, err)
	}

	if err := r.db.Delete(&event).Error; err != nil {
		return fmt.Errorf("Error deleting event with ID %s: %v", id, err)
	}

	return nil
}
