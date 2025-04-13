package buddyservice

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	internalEventService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/internal-event-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type BuddyServiceImpl struct {
	db                   *gorm.DB
	buddyRepository      repository.BuddyRepository
	internalEventService internalEventService.InternalEventService
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

func (b *BuddyServiceImpl) GetBuddyRequests(ctx *gin.Context, uid uuid.UUID) ([]dto.BuddyRequestDTO, error) {
	buddyRequests, err := b.buddyRepository.GetBuddyRequests(uid)
	if err != nil {
		return nil, err
	}
	var buddyRequestDTOs []dto.BuddyRequestDTO
	for _, buddyRequest := range buddyRequests {
		eventDetailDto, err := b.internalEventService.FindById(uid, buddyRequest.EventID)
		var eventDetail dto.EventDetailDTO
		err = mapstructure.Decode(eventDetailDto, &eventDetail)
		if err != nil {
			return nil, err
		}

		buddyRequestDTO, err := mapper.MapBuddyRequestEntityToDto(buddyRequest, &eventDetail)
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
		buddyRequestDTO, err := mapper.MapBuddyRequestEntityToDto(buddyRequest, nil)
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
		buddyRequestDTO, err := mapper.MapBuddyRequestEntityToDto(buddyRequest, nil)
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

func NewBuddyService(db *gorm.DB,
	buddyRepository repository.BuddyRepository,
	internalEventService internalEventService.InternalEventService) BuddyServiceImpl {
	return BuddyServiceImpl{
		db:                   db,
		buddyRepository:      buddyRepository,
		internalEventService: internalEventService,
	}
}
