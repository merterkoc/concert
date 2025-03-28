package eventservice

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"github.com/google/uuid"
)

type eventService struct {
	eventRepo repository.EventRepository
}

func (e *eventService) JoinEvent(ID uuid.UUID, eventId string) error {
	return e.eventRepo.JoinEvent(ID.String(), eventId)
}

func (e *eventService) LeaveEvent(ID uuid.UUID, eventId string) error {
	return e.eventRepo.LeaveEvent(ID.String(), eventId)
}

func (e *eventService) GetEventByUser(ID uuid.UUID) ([]string, error) {
	return e.eventRepo.GetEventByUser(ID.String())
}

func NewEventService(
	eventRepo *repository.EventRepository,
) EventService {
	return &eventService{
		eventRepo: *eventRepo,
	}
}
