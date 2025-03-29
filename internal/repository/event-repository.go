package repository

import (
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) JoinEvent(userID, eventID string) error {
	var UserEvents entity.UserEvents

	err := r.db.Where("user_id = ? AND event_id = ?", userID, eventID).
		FirstOrCreate(&UserEvents, entity.UserEvents{
			UserID:  uuid.MustParse(userID),
			EventID: eventID,
			Status:  "going",
		}).Error

	return err
}

func (r *EventRepository) LeaveEvent(userID, eventID string) error {
	var userEvents entity.UserEvents

	err := r.db.Where("user_id = ? AND event_id = ?", userID, eventID).First(&userEvents).Error
	if err != nil {
		return gorm.ErrRecordNotFound
	}
	err = r.db.Delete(&userEvents).Error
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	return nil
}

func (r *EventRepository) GetEventByUser(id string) ([]string, error) {
	var eventListIDs []string
	err := r.db.Model(&entity.UserEvents{}).
		Where("user_id = ?", id).
		Pluck("event_id", &eventListIDs).Error
	if err != nil {
		return []string{}, fmt.Errorf("failed to get event list: %w", err)
	}
	return eventListIDs, nil
}
