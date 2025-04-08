package external_event_service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	entity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	internalEventService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/internal-event-service"

	"github.com/gin-gonic/gin"
)

type ExternalEventService interface {
	SetInternalService(internalService internalEventService.InternalEventService)
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(ctx *gin.Context, keyword string, location string, page int, size int) ([]dto.EventDTO, error)
	GetEventByIDs(eventIDs []string) ([]entity.EventDetail, error)
}
