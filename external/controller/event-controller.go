package controller

import (
	service "gilab.com/pragmaticreviews/golang-gin-poc/external/event-service"
	entity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
)

type EventController interface {
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(GetEventRequest dto.GetEventRequest) ([]entity.Event, error)
}

type eventController struct {
	eventService service.EventService
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
// @Tags ticketmaster-event
// @Security AccessToken[admin, user]
func (c eventController) FindById(id string) (entity.EventDetail, error) {
	return c.eventService.FindById(id)
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
// @Tags ticketmaster-event
// @Security AccessToken[admin, user]
func (c eventController) FindByKeywordOrLocation(GetEventRequest dto.GetEventRequest) ([]entity.Event, error) {
	return c.eventService.FindByKeywordOrLocation(
		GetEventRequest.Keyword,
		GetEventRequest.Location,
		GetEventRequest.Page,
		GetEventRequest.Size,
	)
}

func NewEventController(eventService service.EventService) EventController {
	return eventController{eventService: eventService}
}
