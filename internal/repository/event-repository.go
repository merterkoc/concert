package repository

import (
	"fmt"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository struct {
	db                 *gorm.DB
	identityRepository *IdentityRepository
}

func NewEventRepository(db *gorm.DB, identityRepository *IdentityRepository) *EventRepository {
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

func (r *EventRepository) GetUsersAvatarByEventId(id string) ([]*string, error) {

	var userEvents []entity.UserEvents

	err := r.db.
		Preload("User").
		Where("event_id = ?", id).
		Find(&userEvents).Error

	if err != nil {
		return []*string{}, fmt.Errorf("failed to get event list: %w", err)
	}

	var userImages []*string
	for _, ue := range userEvents {
		userImages = append(userImages, ue.User.UserImage)
	}

	//var retrievedCourse Course
	//db.Preload("Students").First(&retrievedCourse, course2.ID)
	//fmt.Println("Kurs:", retrievedCourse.Name)
	//for _, student := range retrievedCourse.Students {
	//	fmt.Println("Öğrenci:", student.Name)
	//}

	return userImages, nil
}
