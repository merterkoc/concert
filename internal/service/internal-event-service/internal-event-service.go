package internal_event_service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InternalEventService interface {
	FindById(gin *gin.Context, id string)
	FindByKeywordOrLocation(gin *gin.Context, keyword string, location string, page int, size int)
	JoinEvent(ID uuid.UUID, eventID string) error
	LeaveEvent(ID uuid.UUID, eventID string) error
	GetEventDTOByUserID(gin *gin.Context, ID uuid.UUID)
	GetUsersAvatarByEventId(id string) ([]*string, error)
	GetUsersAvatarByEventIdAndUserId(id string, userID uuid.UUID) ([]*string, error)
	GetEventsByEventIDs(ID uuid.UUID) ([]entity.EventDetail, error)
}
