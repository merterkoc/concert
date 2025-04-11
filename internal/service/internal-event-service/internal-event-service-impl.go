package internal_event_service

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	externalEventService "gilab.com/pragmaticreviews/golang-gin-poc/external/external-event-service"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/mapper"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type internalEventService struct {
	eventRepo            repository.EventRepository
	externalEventService externalEventService.ExternalEventService
}

// FindById implements InternalEventService.
func (e *internalEventService) FindById(c *gin.Context, id string) {
	res, err := externalEventService.NewEventService().FindById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": res})
	return
}

// FindByKeywordOrLocation implements InternalEventService.
func (e *internalEventService) FindByKeywordOrLocation(c *gin.Context, keyword string, location string, page int, size int) {
	res, err := e.externalEventService.FindByKeywordOrLocation(c, keyword, location, page, size)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(res) > 0 {
		uid, exists := c.Get("user_id")
		if !exists {
			c.JSON(500, gin.H{"error": "user_id not found in context"})
			return
		}
		id, err := uuid.Parse(uid.(string))
		if err != nil {
			c.JSON(500, gin.H{"error": "invalid user_id"})
			return
		}

		userEvents, err := e.GetEventIDsByUser(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		joinedEventMap := make(map[string]bool)
		for _, eventID := range userEvents {
			joinedEventMap[eventID] = true
		}

		var events []dto.EventDTO
		for _, event := range res {
			isJoined := joinedEventMap[event.ID]

			participant, err := e.GetUsersAvatarByEventId(event.ID)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}

			eventDto, err := mapper.MapEventEntityToDTO(event, isJoined, participant)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
			}

			events = append(events, *eventDto)
		}
		c.JSON(200, gin.H{"data": events})
	}
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
	eventList, err := e.eventRepo.GetEventIDsByUser(ID.String())
	if err != nil {
		return nil, err
	}
	if len(eventList) == 0 {
		return nil, nil
	}
	return eventList, nil
}

func (e *internalEventService) GetUsersAvatarByEventId(id string) ([]*string, error) {
	return e.eventRepo.GetUsersAvatarByEventId(id)
}

func (e *internalEventService) GetUsersAvatarByEventIdAndUserId(id string, userID uuid.UUID) ([]*string, error) {
	return e.eventRepo.GetUsersAvatarByEventIdAndUserId(id, userID)
}

func (e *internalEventService) GetEventDTOByUserID(c *gin.Context, ID uuid.UUID) {
	eventList, err := e.eventRepo.GetEventIDsByUser(ID.String())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if len(eventList) == 0 {
		c.JSON(200, gin.H{"data": []dto.EventDTO{}})
		return
	}
	events, err := e.externalEventService.GetEventByIDs(eventList)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var eventDTOs []dto.EventDetailDTO

	for _, event := range events {
		participant, err := e.GetUsersAvatarByEventIdAndUserId(event.ID, ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		eventDto, err := mapper.MapEventDetailEntityToEventDetailDto(event, false, participant)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		eventDTOs = append(eventDTOs, *eventDto)
	}
	c.JSON(200, gin.H{"data": eventDTOs})
}

func NewEventService(
	eventRepo repository.EventRepository,
	externalEventService externalEventService.ExternalEventService,
) InternalEventService {
	return &internalEventService{
		eventRepo:            eventRepo,
		externalEventService: externalEventService,
	}
}
