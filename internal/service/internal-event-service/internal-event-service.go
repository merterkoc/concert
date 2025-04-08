package internal_event_service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	"github.com/google/uuid"
)

type InternalEventService interface {
	JoinEvent(ID uuid.UUID, eventID string) error
	LeaveEvent(ID uuid.UUID, eventID string) error
	GetEventIDsByUser(ID uuid.UUID) ([]string, error)
	GetUsersAvatarByEventId(id string) ([]*string, error)
	GetEventsByEventIDs(ID uuid.UUID) ([]entity.EventDetail, error)
}
