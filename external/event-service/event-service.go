package eventservice

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	entity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	"github.com/gin-gonic/gin"
)

type EventService interface {
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(ctx *gin.Context, keyword string, location string, page int, size int) ([]dto.EventDTO, error)
}
