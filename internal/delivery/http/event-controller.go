package controller

import (
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/event/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/service"
)

type EventController interface {
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(GetEventRequest dto.GetEventRequest) ([]entity.Event, error)
}

type controller struct {
	service service.EventService
}

// FindById is a controller method
// that returns an event by id
// @Summary Get event by id
// @Description Get event by id
// @ID get-event-by-id
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} entity.EventDetail
// @Router /events/{id} [get]
// @Tags events
func (c controller) FindById(id string) (entity.EventDetail, error) {
	return c.service.FindById(id)
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
func (c controller) FindByKeywordOrLocation(GetEventRequest dto.GetEventRequest) ([]entity.Event, error) {
	return c.service.FindByKeywordOrLocation(
		GetEventRequest.Keyword,
		GetEventRequest.Location,
		GetEventRequest.Page,
		GetEventRequest.Size,
	)
}

func New(ss service.EventService) EventController {
	return controller{service: ss}
}
