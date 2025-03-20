package controller

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/entity"
	"gilab.com/pragmaticreviews/golang-gin-poc/service"
)

type EventController interface {
	FindByKeyword(keyword string) []entity.Event
	FindByLocation(location string) []entity.Event
}

type controller struct {
	service service.EventService
}

func (c controller) FindByKeyword(keyword string) []entity.Event {
	return c.service.FindByKeyword(keyword)
}

func (c controller) FindByLocation(location string) []entity.Event {
	return c.service.FindByLocation(location)
}

func New(ss service.EventService) EventController {
	return controller{service: ss}
}
