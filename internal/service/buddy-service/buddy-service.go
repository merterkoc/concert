package buddyservice

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BuddyService interface {
	GetBuddyRequests(ctx *gin.Context, uid uuid.UUID) ([]dto.BuddyRequestDTO, error)
	CreateBuddyRequest(senderID uuid.UUID, receiverID uuid.UUID, eventID uuid.UUID) error
	AcceptBuddyRequest(senderID uuid.UUID, receiverID uuid.UUID, eventID uuid.UUID) error
	RejectBuddyRequest(senderID uuid.UUID, receiverID uuid.UUID, eventID uuid.UUID) error
	BlockBuddyRequest(senderID uuid.UUID, receiverID uuid.UUID, eventID uuid.UUID) error
	GetBuddyRequestsByEventID(eventID uuid.UUID) ([]dto.BuddyRequestDTO, error)
	GetBuddyRequestsByUserID(userID uuid.UUID) ([]dto.BuddyRequestDTO, error)
}
