package repository

import (
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
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

func (r *EventRepository) GetEventIDsByUser(id string) ([]string, error) {
	var eventListIDs []string
	err := r.db.Model(&entity.UserEvents{}).
		Where("user_id = ?", id).
		Pluck("event_id", &eventListIDs).Error
	if err != nil {
		return []string{}, fmt.Errorf("failed to get event list: %w", err)
	}
	return eventListIDs, nil
}

func (r *EventRepository) GetUsersAvatarByEventId(id string) ([]dto.ParticipantsAvatar, error) {

	var userEvents []entity.UserEvents

	err := r.db.
		Preload("User").
		Where("event_id = ?", id).
		Find(&userEvents).Error

	if err != nil {
		return []dto.ParticipantsAvatar{}, fmt.Errorf("failed to get event list: %w", err)
	}

	var userImages []dto.ParticipantsAvatar
	for _, ue := range userEvents {
		userImages = append(userImages, dto.ParticipantsAvatar{
			ID:        ue.User.ID,
			UserImage: ue.User.UserImage,
		})
	}

	return userImages, nil
}

func (r *EventRepository) GetUsersAvatarByEventIdAndUserId(id string, userID uuid.UUID) ([]dto.ParticipantsAvatar, error) {

	var buddyship []entity.Buddyship

	err := r.db.
		Preload("User1").
		Preload("User2").
		Where("event_id = ? AND (user1_id = ? OR user2_id = ?)", id, userID, userID).
		Find(&buddyship).Error

	if err != nil {
		return []dto.ParticipantsAvatar{}, fmt.Errorf("failed to get event list: %w", err)
	}

	var userImages []dto.ParticipantsAvatar
	for _, ue := range buddyship {
		userImages = append(userImages, dto.ParticipantsAvatar{
			ID:        ue.User1.ID,
			UserImage: ue.User1.UserImage,
		})
	}
	for _, ue := range buddyship {
		userImages = append(userImages, dto.ParticipantsAvatar{
			ID:        ue.User2.ID,
			UserImage: ue.User2.UserImage,
		})
	}

	return userImages, nil
}

func (r *EventRepository) IsJoined(eventID string, uid uuid.UUID) (bool, error) {
	var userEvents []entity.UserEvents
	err := r.db.Where("event_id = ? AND user_id = ?", eventID, uid).Find(&userEvents).Error
	if err != nil {
		return false, err
	}
	return len(userEvents) > 0, nil
}
