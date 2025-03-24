package controller

import (
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/service"
)

type EventController interface {
	FindByKeywordOrLocation(keyword string, location string, page int, size int) (entity.Event, error)
}

type controller struct {
	service service.EventService
}

// FindByKeywordOrLocation is a controller method
// that returns an event by keyword
// @Summary Get event by keyword
// @Description Get event by keyword
// @ID get-event-by-keyword
// @Produce json
// @Param keyword query string false "Keyword"
// @Param location query string false "Location"
// @Param page query int 1 "Page"
// @Param size query int 3 "Size"
// @Success 200 {object} entity.Event
// @Router /events [get]
// @Tags events
func (c controller) FindByKeywordOrLocation(keyword string, location string, page int, size int) (entity.Event, error) {
	return c.service.FindByKeywordOrLocation(keyword, location, page, size)
}

func New(ss service.EventService) EventController {
	return controller{service: ss}
}
