package internal_event_service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InternalEventService interface {
	FindById(uuid uuid.UUID, id string) (dto.EventDetailDTO, error)
	FindByKeywordOrLocation(gin *gin.Context, keyword string, location string, page int, size int)
	JoinEvent(ID uuid.UUID, eventID string) error
	LeaveEvent(ID uuid.UUID, eventID string) error
	GetEventDTOByUserID(gin *gin.Context, ID uuid.UUID)
	GetUsersAvatarByEventId(id string) ([]dto.ParticipantsAvatar, error)
	GetUsersAvatarByEventIdAndUserId(id string, userID uuid.UUID) ([]dto.ParticipantsAvatar, error)
	GetEventsByEventIDs(ID uuid.UUID) ([]entity.EventDetail, error)
}
