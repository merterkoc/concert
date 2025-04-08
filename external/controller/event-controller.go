package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	externalEventService "gilab.com/pragmaticreviews/golang-gin-poc/external/external-event-service"
	"github.com/gin-gonic/gin"
)

type EventController interface {
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(c *gin.Context, GetEventRequest dto.GetEventRequest) ([]dto.EventDTO, error)
}

type eventController struct {
	externalEventService externalEventService.ExternalEventService
}

// FindById is a controller method
// that returns an event by id
// @Summary Get event by id
// @Description Get event by id
// @ID get-event-by-id
// @Produce json
// @Param id path string true "Id"
// @Success 200 {object} entity.EventDetail "Return event detail successfully"
// @Router /events/{id} [get]
// @Tags ticketmaster-event
// @Security AccessToken[admin, user]
func (c eventController) FindById(id string) (entity.EventDetail, error) {
	return c.externalEventService.FindById(id)
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
// @Success 200 {object} dto.EventDTO "Return event successfully"
// @Router /events [get]
// @Tags ticketmaster-event
// @Security AccessToken[admin, user]
func (c eventController) FindByKeywordOrLocation(ctx *gin.Context, GetEventRequest dto.GetEventRequest) ([]dto.EventDTO, error) {
	return c.externalEventService.FindByKeywordOrLocation(
		ctx,
		GetEventRequest.Keyword,
		GetEventRequest.Location,
		GetEventRequest.Page,
		GetEventRequest.Size,
	)
}

func NewEventController(eventService externalEventService.ExternalEventService) EventController {
	return eventController{externalEventService: eventService}
}
