package eventservice

import (
	"encoding/json"
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/mapper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"

	internalEventService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/event-service"
)

type eventService struct {
	internalEventService internalEventService.EventService
	apiURL               string
}

func (e *eventService) FindById(id string) (entity.EventDetail, error) {
	baseURL, err := url.Parse(e.apiURL + "/events/" + id + ".json")
	params := url.Values{}
	params.Add("apikey", envService.GetEnvServiceInstance().Env.TicketMasterAPIToken)
	baseURL.RawQuery = params.Encode()
	if err != nil {
		fmt.Println("Error parsing URL")
		return entity.EventDetail{}, fmt.Errorf("invalid API URL: %w", err)
	}
	fmt.Println("Request URL: " + baseURL.String())

	// HTTP GET isteği yap
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return entity.EventDetail{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return entity.EventDetail{}, fmt.Errorf("unexpected API response: %d %s, body: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
	}

	if err != nil {
		return entity.EventDetail{}, fmt.Errorf("failed to read API response: %w", err)
	}

	var response entity.EventDetail
	err = json.Unmarshal(body, &response)
	if err != nil {
		return entity.EventDetail{}, fmt.Errorf("JSON parse error: %w", err)
	}

	return response, nil

}

func (e *eventService) FindByKeywordOrLocation(c *gin.Context, keyword string, location string, page int, size int) ([]dto.EventDTO, error) {
	baseURL, err := url.Parse(e.apiURL + "/events.json")
	if err != nil {
		return []dto.EventDTO{}, fmt.Errorf("invalid API URL: %w", err)
	}

	params := url.Values{}

	if location != "" {
		params.Add("latlong", strings.ReplaceAll(location, " ", ""))
	}
	if keyword != "" {
		params.Add("keyword", keyword)
	}
	if keyword != "" {
		params.Add("keyword", keyword)
	}
	if size <= 0 {
		size = 100
	}
	params.Add("size", strconv.Itoa(size))
	if page > 0 {
		params.Add("page", strconv.Itoa(page))
	}
	params.Add("apikey", envService.GetEnvServiceInstance().Env.TicketMasterAPIToken)

	baseURL.RawQuery = params.Encode()

	fmt.Println("Request URL: " + baseURL.String())

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return []dto.EventDTO{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return []dto.EventDTO{}, fmt.Errorf("unexpected API response: %d %s, body: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
	}

	if err != nil {
		return []dto.EventDTO{}, fmt.Errorf("failed to read API response: %w", err)
	}

	var response entity.Embedde
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []dto.EventDTO{}, fmt.Errorf("JSON parse error: %w", err)
	}

	if len(response.ApiResponse.Events) > 0 {
		uid, exists := c.Get("user_id")
		if !exists {
			return []dto.EventDTO{}, fmt.Errorf("user id not found")
		}
		id, err := uuid.Parse(uid.(string))
		if err != nil {
			return []dto.EventDTO{}, fmt.Errorf("failed to parse user id: %w", err)
		}

		userEvents, err := e.internalEventService.GetEventByUser(id)
		if err != nil {
			return nil, err
		}

		joinedEventMap := make(map[string]bool)
		for _, eventID := range userEvents {
			joinedEventMap[eventID] = true
		}

		var events []dto.EventDTO
		for _, event := range response.ApiResponse.Events {
			isJoined := joinedEventMap[event.ID]

			participant, err := e.internalEventService.GetUsersAvatarByEventId(event.ID)
			if err != nil {
				return nil, err
			}

			eventDto, err := mapper.MapEventEntityToDTO(event, isJoined, participant)
			if err != nil {
				return nil, fmt.Errorf("failed to map event entity to dto: %w", err)
			}

			events = append(events, *eventDto)
		}
		return events, nil
	}

	return []dto.EventDTO{}, fmt.Errorf("no events found")
}

func NewEventService(internalEventService internalEventService.EventService) EventService {
	return &eventService{
		apiURL:               "https://app.ticketmaster.com/discovery/v2",
		internalEventService: internalEventService,
	}
}
