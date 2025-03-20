package service

import "gilab.com/pragmaticreviews/golang-gin-poc/entity"

type EventService interface {
	FindByKeyword(keyword string) []entity.Event
	FindByLocation(location string) []entity.Event
}

type eventService struct {
	events []entity.Event
}

func (e eventService) FindByKeyword(keyword string) []entity.Event {
	return e.events
}

func (e eventService) FindByLocation(location string) []entity.Event {
	return e.events
}

func NewEventService() EventService {
	return &eventService{
		events: []entity.Event{
			{ID: 1, Name: "Event 1", Location: "Location 1", Date: "2021-07-01"},
			{ID: 2, Name: "Event 2", Location: "Location 2", Date: "2021-07-02"},
			{ID: 3, Name: "Event 3", Location: "Location 3", Date: "2021-07-03"},
		},
	}
}
