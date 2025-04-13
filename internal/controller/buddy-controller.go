package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	buddyservice "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/buddy-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BuddyController interface {
	GetBuddyRequests(ctx *gin.Context, uid uuid.UUID) ([]dto.BuddyRequestDTO, error)
	CreateBuddyRequest(uid uuid.UUID, dto dto.CreateBuddyRequestDTO) error
	AcceptBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error
	RejectBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error
	BlockBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error
	GetBuddyRequestsByEventID(eventID string) ([]dto.BuddyRequestDTO, error)
}

type buddyController struct {
	buddyService buddyservice.BuddyServiceImpl
}

// AcceptBuddyRequest implements BuddyController.
// @Summary Accept Buddy Request
// @Description Accept a buddy request between users for a specific event. Only the receiver can accept the request.
// @Accept  json
// @Produce json
// @Param id path string true "Buddy Request ID"
// @Success 200
// @Router /buddy/requests/{id}/accept [post]
// @Tags buddy
// @Security AccessToken[admin, user]
func (b *buddyController) AcceptBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	err := b.buddyService.AcceptBuddyRequest(uuid, buddyRequestID)
	if err != nil {
		return err
	}
	return nil
}

// RejectBuddyRequest implements BuddyController.
// @Summary Reject Buddy Request
// @Description Reject Buddy Request
// @ID reject-buddy-request
// @Accept  json
// @Produce json
// @Param id path string true "Buddy Request ID"
// @Success 200
// @Router /buddy/requests/{id}/reject [post]
// @Tags buddy
// @Security AccessToken[admin, user]
func (b *buddyController) RejectBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	err := b.buddyService.RejectBuddyRequest(uuid, buddyRequestID)
	if err != nil {
		return err
	}
	return nil
}

// BlockBuddyRequest implements BuddyController.
// @Summary Block Buddy Request
// @Description Block Buddy Request
// @ID block-buddy-request
// @Accept  json
// @Produce json
// @Param id path string true "Buddy Request ID"
// @Success 200 {object} entity.BuddyRequest "Return buddy request successfully"
// @Router /buddy/requests/{id}/block [post]
// @Tags buddy
// @Security AccessToken[admin, user]
func (b *buddyController) BlockBuddyRequest(uuid uuid.UUID, buddyRequestID uuid.UUID) error {
	err := b.buddyService.BlockBuddyRequest(uuid, buddyRequestID)
	if err != nil {
		return err
	}
	return nil
}

// CreateBuddyRequest implements BuddyController.
// @Summary Create Buddy Request
// @Description Create a buddy request between users for a specific event
// @Accept  json
// @Produce json
// @Param dto body dto.CreateBuddyRequestDTO true "Buddy Request Information"
// @Success 200
// @Router /buddy/requests [post]
// @Tags buddy
// @Security AccessToken[admin, user]
func (b *buddyController) CreateBuddyRequest(uid uuid.UUID, dto dto.CreateBuddyRequestDTO) error {
	receiverID := uuid.MustParse(dto.ReceiverID)

	err := b.buddyService.CreateBuddyRequest(uid, receiverID, dto.EventID)
	if err != nil {
		return err
	}
	return nil
}

func (b *buddyController) GetBuddyRequestsByEventID(eventID string) ([]dto.BuddyRequestDTO, error) {
	res, err := b.buddyService.GetBuddyRequestsByEventID(eventID)
	if err != nil {
		return res, err
	}
	return res, nil
}

func NewBuddyController(buddyService buddyservice.BuddyServiceImpl) BuddyController {
	return &buddyController{
		buddyService: buddyService,
	}
}

// GetBuddyRequests is a controller method
// @Summary Get Buddy Requests
// @Description Retrieve a list of buddy requests for a user
// @Produce json
// @Success 200 {array} dto.BuddyRequestDTO "Return buddy requests successfully"
// @Router /buddy/requests [get]
// @Tags buddy
// @Security AccessToken[admin, user]
func (b *buddyController) GetBuddyRequests(ctx *gin.Context, uid uuid.UUID) ([]dto.BuddyRequestDTO, error) {
	res, err := b.buddyService.GetBuddyRequests(ctx, uid)
	if err != nil {
		return res, err
	}
	return res, nil
}
