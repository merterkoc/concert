package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"
)

type eventService struct {
	apiURL string
}

func (e *eventService) FindByKeywordOrLocation(keyword string, location string, page int, size int) (entity.Event, error) {
	// Temel API URL'si
	baseURL, err := url.Parse(e.apiURL)
	if err != nil {
		return entity.Event{}, fmt.Errorf("invalid API URL: %w", err)
	}

	// Query parametrelerini oluştur
	params := url.Values{}

	if location != "" {
		params.Add("latlong", location)
	}
	if keyword != "" {
		params.Add("keyword", keyword)
	}
	params.Add("size", strconv.Itoa(size))
	params.Add("page", strconv.Itoa(page))
	params.Add("apikey", envService.GetEnvServiceInstance().Env.TicketMasterAPIToken)

	// Query parametrelerini API URL'sine ekle
	baseURL.RawQuery = params.Encode()

	fmt.Println("Request URL: " + baseURL.String())

	// HTTP GET isteği yap
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return entity.Event{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	// **HTTP durum kodunu kontrol et**
	if resp.StatusCode != http.StatusOK {
		return entity.Event{}, fmt.Errorf("unexpected API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	// Yanıtı oku
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.Event{}, fmt.Errorf("failed to read API response: %w", err)
	}

	// JSON yanıtını çözümle
	var response entity.Embedde
	err = json.Unmarshal(body, &response)
	if err != nil {
		return entity.Event{}, fmt.Errorf("JSON parse error: %w", err)
	}

	// Eğer veri varsa, ilk etkinliği döndür
	if len(response.ApiResponse.Events) > 0 {
		return response.ApiResponse.Events[0], nil
	}

	return entity.Event{}, fmt.Errorf("no events found")
}

func NewEventService(
	envService *envService.EnvService,
) EventService {
	return &eventService{
		apiURL: "https://app.ticketmaster.com/discovery/v2/events.json",
	}
}

// Helper function to check if a string contains a keyword (case-insensitive)
func contains(text, keyword string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(keyword))
}
