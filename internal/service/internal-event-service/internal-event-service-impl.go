package internal_event_service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	externalEventService "gilab.com/pragmaticreviews/golang-gin-poc/external/external-event-service"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"github.com/google/uuid"
)

type internalEventService struct {
	eventRepo repository.EventRepository

	externalEventService externalEventService.ExternalEventService
}

func (e *internalEventService) GetEventsByEventIDs(ID uuid.UUID) ([]entity.EventDetail, error) {
	userEventListIDs, err := e.eventRepo.GetEventIDsByUser(ID.String())
	if err != nil {
		return nil, err
	}
	events, err := e.externalEventService.GetEventByIDs(userEventListIDs)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (e *internalEventService) JoinEvent(ID uuid.UUID, eventId string) error {
	return e.eventRepo.JoinEvent(ID.String(), eventId)
}

func (e *internalEventService) LeaveEvent(ID uuid.UUID, eventId string) error {
	return e.eventRepo.LeaveEvent(ID.String(), eventId)
}

func (e *internalEventService) GetEventIDsByUser(ID uuid.UUID) ([]string, error) {
	return e.eventRepo.GetEventIDsByUser(ID.String())
}

func (e *internalEventService) GetUsersAvatarByEventId(id string) ([]*string, error) {
	return e.eventRepo.GetUsersAvatarByEventId(id)
}

func NewEventService(
	eventRepo *repository.EventRepository,
	externalEventService externalEventService.ExternalEventService,
) InternalEventService {
	return &internalEventService{
		eventRepo: *eventRepo,
	}
}
