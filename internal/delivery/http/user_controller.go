package controller

import (
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/service"
)

type UserController interface {
	CreateUser(name string) (entity.User, error)
}

type controller struct {
	service service.UserService
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
func (c controller) CreateUser(id string) (entity.UserService, error) {
	return c.service.CreateUser(id)
}

func New(ss service.UserService) EventController {
	return controller{service: ss}
}
