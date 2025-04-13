package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	internalEventService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/internal-event-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type EventController interface {
	FindById(cc *gin.Context, uuid uuid.UUID, id string)
	FindByKeywordOrLocation(c *gin.Context, GetEventRequest dto.GetEventRequest)
	JoinEvent(ID uuid.UUID, eventID string) error
	LeaveEvent(ID uuid.UUID, eventID string) error
	GetEventByUser(c *gin.Context, ID uuid.UUID)
	GetEventByUserID(c *gin.Context, ID uuid.UUID)
}

type eventController struct {
	eventService internalEventService.InternalEventService
}

// FindById is a controller method
// that returns an event by id
// @Summary Get event by id
// @Description Get event by id
// @ID get-event-by-id
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} entity.EventDetail "Return event detail successfully"
// @Router /events/{id} [get]
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) FindById(cc *gin.Context, uuid uuid.UUID, id string) {
	byId, err := c.eventService.FindById(uuid, id)
	if err != nil {
		return
	}

	cc.JSON(http.StatusOK, gin.H{
		"data": byId,
	})
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
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) FindByKeywordOrLocation(ctx *gin.Context, GetEventRequest dto.GetEventRequest) {
	c.eventService.FindByKeywordOrLocation(
		ctx,
		GetEventRequest.Keyword,
		GetEventRequest.Location,
		GetEventRequest.Page,
		GetEventRequest.Size,
	)
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
//	@ID https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#InvalidIfaceAssignget-event-by-user
//
// @Produce json
//
//	@Success 200 {object} entity.EventDetail "Return event detail successfully"
//	@Router /events/user [get]
//
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) GetEventByUser(gin *gin.Context, ID uuid.UUID) {
	c.eventService.GetEventDTOByUserID(gin, ID)

}

// GetEventByUserID is a controller method
// that returns an event by user id
// @Summary Get event by user id
// @Description Get event by user id
// @ID get-event-by-user-id
// @Produce json
// @Param userId path string true "User ID"
// @Success 200 {object} entity.EventDetail
// @Router /events/user/{userId} [get]
// @Tags events
// @Security AccessToken[admin, user]
func (c eventController) GetEventByUserID(gin *gin.Context, ID uuid.UUID) {
	c.eventService.GetEventDTOByUserID(gin, ID)
}

func NewEventController(eventService internalEventService.InternalEventService) EventController {
	return eventController{eventService: eventService}
}
