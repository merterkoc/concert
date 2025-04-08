package external_event_service

import (
	entity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	"github.com/gin-gonic/gin"
)

type ExternalEventService interface {
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(ctx *gin.Context, keyword string, location string, page int, size int) ([]entity.Event, error)
	GetEventByIDs(eventIDs []string) ([]entity.EventDetail, error)
}
