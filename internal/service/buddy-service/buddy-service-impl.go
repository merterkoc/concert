package buddyservice

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BuddyServiceImpl struct {
	db              *gorm.DB
	buddyRepository repository.BuddyRepository
}

func (b *BuddyServiceImpl) AcceptBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	b.buddyRepository.AcceptBuddyRequest(uuid, buddyRequestID)
	if err := b.db.Error; err != nil {
		return err
	}
	return nil
}

func (b *BuddyServiceImpl) BlockBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	b.buddyRepository.BlockBuddyRequest(uuid, buddyRequestID)
	if err := b.db.Error; err != nil {
		return err
	}
	return nil
}

func (b *BuddyServiceImpl) CreateBuddyRequest(senderID uuid.UUID, receiverID uuid.UUID, eventID string) error {
	b.buddyRepository.CreateBuddyRequest(senderID, receiverID, eventID)
	if err := b.db.Error; err != nil {
		return err
	}
	return nil
}

func (b *BuddyServiceImpl) GetBuddyRequests(uid uuid.UUID) ([]dto.BuddyRequestDTO, error) {
	buddyRequests, err := b.buddyRepository.GetBuddyRequests(uid)
	if err != nil {
		return nil, err
	}
	var buddyRequestDTOs []dto.BuddyRequestDTO
	for _, buddyRequest := range buddyRequests {
		buddyRequestDTO, err := mapper.MapBuddyRequestEntityToDto(buddyRequest)
		if err != nil {
			return nil, err
		}
		buddyRequestDTOs = append(buddyRequestDTOs, *buddyRequestDTO)
	}

	return buddyRequestDTOs, nil
}

func (b *BuddyServiceImpl) GetBuddyRequestsByEventID(eventID string) ([]dto.BuddyRequestDTO, error) {
	buddyRequests, err := b.buddyRepository.GetBuddyRequestsByEventID(eventID)
	if err != nil {
		return nil, err
	}
	var buddyRequestDTOs []dto.BuddyRequestDTO
	for _, buddyRequest := range buddyRequests {
		buddyRequestDTO, err := mapper.MapBuddyRequestEntityToDto(buddyRequest)
		if err != nil {
			return nil, err
		}
		buddyRequestDTOs = append(buddyRequestDTOs, *buddyRequestDTO)
	}
	return buddyRequestDTOs, nil
}

func (b *BuddyServiceImpl) GetBuddyRequestsByUserID(userID uuid.UUID) ([]dto.BuddyRequestDTO, error) {
	buddyRequests, err := b.buddyRepository.GetBuddyRequestsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var buddyRequestDTOs []dto.BuddyRequestDTO
	for _, buddyRequest := range buddyRequests {
		buddyRequestDTO, err := mapper.MapBuddyRequestEntityToDto(buddyRequest)
		if err != nil {
			return nil, err
		}
		buddyRequestDTOs = append(buddyRequestDTOs, *buddyRequestDTO)
	}
	return buddyRequestDTOs, nil
}

func (b *BuddyServiceImpl) RejectBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	b.buddyRepository.RejectBuddyRequest(uuid, buddyRequestID)
	if err := b.db.Error; err != nil {
		return err
	}
	return nil
}

func NewBuddyService(db *gorm.DB, buddyRepository repository.BuddyRepository) BuddyServiceImpl {
	return BuddyServiceImpl{
		db:              db,
		buddyRepository: buddyRepository,
	}
}
