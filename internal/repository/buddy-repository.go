package repository

import (
	"fmt"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BuddyRepository interface {
	GetBuddyRequests(uid uuid.UUID) ([]entity.BuddyRequest, error)
	CreateBuddyRequest(buddyRequestID uuid.UUID, receiverID uuid.UUID, eventID string) error
	AcceptBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error
	RejectBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error
	BlockBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error
	GetBuddyRequestsByEventID(eventID string) ([]entity.BuddyRequest, error)
	GetBuddyRequestsByUserID(userID uuid.UUID) ([]entity.BuddyRequest, error)
}

func NewBuddyRepository(db *gorm.DB) BuddyRepository {
	return &BuddyRepositoryImpl{db: db}
}

type BuddyRepositoryImpl struct {
	db *gorm.DB
}

// AcceptBuddyRequest implements BuddyRepository.
func (r *BuddyRepositoryImpl) AcceptBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	r.db.Model(&entity.BuddyRequest{}).
		Where("id = ? AND receiver_id = ?", buddyRequestID, uuid).
		Updates(entity.BuddyRequest{Status: entity.Accepted})
	if err := r.db.Error; err != nil {
		return err
	}
	var buddyRequest entity.BuddyRequest

	if err := r.db.Model(&entity.BuddyRequest{}).
		Where("id = ?", buddyRequestID).
		First(&buddyRequest).Error; err != nil {
		return fmt.Errorf("error retrieving updated buddy request: %w", err)
	}

	entityBuddyshipt := entity.Buddyship{
		User1ID: buddyRequest.SenderID,
		User2ID: buddyRequest.ReceiverID,
		EventID: buddyRequest.EventID,
	}
	r.db.Model(&entity.Buddyship{}).
		Create(&entityBuddyshipt)
	if err := r.db.Error; err != nil {
		return err
	}
	return nil
}

// BlockBuddyRequest implements BuddyRepository.
func (r *BuddyRepositoryImpl) BlockBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	r.db.Model(&entity.BuddyRequest{}).
		Where("id = ? AND receiver_id = ?", buddyRequestID, uuid).
		Updates(entity.BuddyRequest{Status: entity.Blocked})
	if err := r.db.Error; err != nil {
		return err
	}
	return nil
}

// CreateBuddyRequest implements BuddyRepository.
func (r *BuddyRepositoryImpl) CreateBuddyRequest(senderID uuid.UUID, receiverID uuid.UUID, eventID string) error {
	buddyRequestEntity := &entity.BuddyRequest{
		SenderID:   senderID.String(),
		ReceiverID: receiverID.String(),
		EventID:    eventID,
		Status:     entity.Pending,
	}
	r.db.
		Model(&entity.BuddyRequest{}).
		Create(buddyRequestEntity)

	return nil
}

// GetBuddyRequestsByEventID implements BuddyRepository.
func (r *BuddyRepositoryImpl) GetBuddyRequestsByEventID(eventID string) ([]entity.BuddyRequest, error) {
	r.db.Model(&entity.BuddyRequest{}).
		Where("event_id = ?", eventID).
		Preload("Sender").
		Preload("Receiver").
		Find(&[]entity.BuddyRequest{})
	if err := r.db.Error; err != nil {
		return nil, err
	}
	var buddyRequests []entity.BuddyRequest
	err := r.db.Where("event_id = ?", eventID).Find(&buddyRequests).Error
	if err != nil {
		return nil, err
	}
	return buddyRequests, nil
}

// GetBuddyRequestsByUserID implements BuddyRepository.
func (r *BuddyRepositoryImpl) GetBuddyRequestsByUserID(userID uuid.UUID) ([]entity.BuddyRequest, error) {
	r.db.Model(&entity.BuddyRequest{}).
		Where("receiver_id = ?", userID).
		Preload("Sender").
		Preload("Receiver").
		Find(&[]entity.BuddyRequest{})
	if err := r.db.Error; err != nil {
		return nil, err
	}
	var buddyRequests []entity.BuddyRequest
	err := r.db.Where("sender_id = ? OR receiver_id = ?", userID, userID).Find(&buddyRequests).Error
	if err != nil {
		return nil, err
	}
	return buddyRequests, nil
}

// RejectBuddyRequest implements BuddyRepository.
func (r *BuddyRepositoryImpl) RejectBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	r.db.Model(&entity.BuddyRequest{}).
		Where("id = ? AND receiver_id = ?", buddyRequestID, uuid).
		Updates(entity.BuddyRequest{Status: entity.Rejected})
	if err := r.db.Error; err != nil {
		return err
	}
	return nil
}

func (r *BuddyRepositoryImpl) GetBuddyRequests(uid uuid.UUID) ([]entity.BuddyRequest, error) {
	var buddyRequests []entity.BuddyRequest
	err := r.db.Preload("Sender").Preload("Receiver").Where("sender_id = ? OR receiver_id = ?", uid, uid).Find(&buddyRequests).Error
	if err != nil {
		return nil, err
	}
	return buddyRequests, nil
}
