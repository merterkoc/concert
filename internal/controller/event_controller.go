package controller

import (
	internalEventService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/internal-event-service"
	"github.com/google/uuid"
)

type EventController interface {
	JoinEvent(ID uuid.UUID, eventID string) error
	LeaveEvent(ID uuid.UUID, eventID string) error
	GetEventByUser(ID uuid.UUID)
}

type eventController struct {
	eventService internalEventService.InternalEventService
}

// JoinEvent is a controller method
// that joins an event
// @Summary Join event
// @Description Join event
// @ID join-event
// @Produce json
// @Param eventId path string true "Event ID"
// @Success 200 {object} entity.EventDetail "Return event detail successfully"
// @Router /events/{eventId}/join [post]
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) JoinEvent(ID uuid.UUID, eventID string) error {
	return c.eventService.JoinEvent(ID, eventID)
}

// LeaveEvent is a controller method
// that leaves an event
// @Summary Leave event
// @Description Leave event
// @ID leave-event
// @Produce json
// @Param eventId path string true "Event ID"
//
// @Success 200 {object} entity.EventDetail "Return event detail successfully"
// @Router /events/{eventId}/leave [post]
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) LeaveEvent(ID uuid.UUID, eventID string) error {
	return c.eventService.LeaveEvent(ID, eventID)
}

// GetEventByUser is a controller method
// that returns an event by user
// @Summary Get event by user
// @Description Get event by user
//
//	@ID get-event-by-user
//
// @Produce json
//
//	@Success 200 {object} entity.EventDetail "Return event detail successfully"
//	@Router /events/user [get]
//
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) GetEventByUser(ID uuid.UUID) {
	c.eventService.GetEventIDsByUser(ID)
}

func NewEventController(eventService internalEventService.InternalEventService) EventController {
	return eventController{eventService: eventService}
}
