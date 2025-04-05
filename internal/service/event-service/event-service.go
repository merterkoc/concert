package eventservice

import "github.com/google/uuid"

type EventService interface {
	JoinEvent(ID uuid.UUID, eventID string) error
	LeaveEvent(ID uuid.UUID, eventID string) error
	GetEventByUser(ID uuid.UUID) ([]string, error)
	GetUsersAvatarByEventId(id string) ([]*string, error)
}
