package mapper

import (
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/entity"

	"gilab.com/pragmaticreviews/golang-gin-poc/external/event/dto"
	"github.com/mitchellh/mapstructure"
)

func MapEventEntityToDTO(event entity.Event, isJoined bool, participantAvatars []dto.ParticipantsAvatar) (*dto.EventDTO, error) {
	var eventDto dto.EventDTO

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           &eventDto,
		ErrorUnused:      false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create decoder: %w", err)
	}

	err = decoder.Decode(event)
	if err != nil {
		return nil, fmt.Errorf("failed to map event entity to dto: %w", err)
	}

	var imageURL []string
	for _, image := range event.Embedded.Attractions[0].Images {
		imageURL = append(imageURL, image.URL)
	}
	eventDto.City = event.Embedded.Venues[0].City.Name
	eventDto.Country = event.Embedded.Venues[0].Country.Name
	eventDto.Locale = event.Locale
	eventDto.VenueName = event.Embedded.Venues[0].Name
	eventDto.Images = imageURL
	eventDto.Start = event.Dates.Start.LocalDate
	eventDto.URL = event.Embedded.Attractions[0].URL
	eventDto.IsJoined = isJoined
	eventDto.ParticipantAvatars = participantAvatars

	return &eventDto, nil
}
