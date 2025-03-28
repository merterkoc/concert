package eventservice

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	entity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/domain"
	envService "gilab.com/pragmaticreviews/golang-gin-poc/internal/config"
)

type eventService struct {
	apiURL string
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
	// **HTTP durum kodunu kontrol et**
	if resp.StatusCode != http.StatusOK {
		return entity.EventDetail{}, fmt.Errorf("unexpected API response: %d %s, body: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
	}

	// Yanıtı oku

	if err != nil {
		return entity.EventDetail{}, fmt.Errorf("failed to read API response: %w", err)
	}

	// JSON yanıtını çözümle
	var response entity.EventDetail
	err = json.Unmarshal(body, &response)
	if err != nil {
		return entity.EventDetail{}, fmt.Errorf("JSON parse error: %w", err)
	}

	return response, nil

}

func (e *eventService) FindByKeywordOrLocation(keyword string, location string, page int, size int) ([]entity.Event, error) {
	// Temel API URL'si
	baseURL, err := url.Parse(e.apiURL + "/events.json")
	if err != nil {
		return []entity.Event{}, fmt.Errorf("invalid API URL: %w", err)
	}

	// Query parametrelerini oluştur
	params := url.Values{}

	if location != "" {
		params.Add("latlong", strings.ReplaceAll(location, " ", ""))
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
		return []entity.Event{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// **HTTP durum kodunu kontrol et**
	if resp.StatusCode != http.StatusOK {
		return []entity.Event{}, fmt.Errorf("unexpected API response: %d %s, body: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(body))
	}

	// Yanıtı oku

	if err != nil {
		return []entity.Event{}, fmt.Errorf("failed to read API response: %w", err)
	}

	// JSON yanıtını çözümle
	var response entity.Embedde
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []entity.Event{}, fmt.Errorf("JSON parse error: %w", err)
	}

	// Eğer veri varsa, ilk etkinliği döndür
	if len(response.ApiResponse.Events) > 0 {
		return response.ApiResponse.Events, nil
	}

	return []entity.Event{}, fmt.Errorf("no events found")
}

func NewEventService(
	envService *envService.EnvService,
) EventService {
	return &eventService{
		apiURL: "https://app.ticketmaster.com/discovery/v2",
	}
}
