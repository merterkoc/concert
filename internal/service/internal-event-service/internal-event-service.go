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
	GetEventIDsByUser(ID uuid.UUID) ([]string, error)
	GetUsersAvatarByEventId(id string) ([]*string, error)
	GetEventsByEventIDs(ID uuid.UUID) ([]entity.EventDetail, error)
}
