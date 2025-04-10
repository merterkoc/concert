package mapper

import (
	"fmt"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"github.com/mitchellh/mapstructure"
)

func MapBuddyRequestEntityToDto(user entity.BuddyRequest) (*dto.BuddyRequestDTO, error) {
	var buddyRequestDTO dto.BuddyRequestDTO

	err := mapstructure.Decode(user, &buddyRequestDTO)
	if err != nil {
		return nil, fmt.Errorf("failed to map user entity to dto: %w", err)
	}

	return &buddyRequestDTO, nil
}

func MapBuddyRequestDtoToEntity(buddyRequestDTO dto.BuddyRequestDTO) (*entity.BuddyRequest, error) {
	var buddyRequest entity.BuddyRequest

	err := mapstructure.Decode(buddyRequestDTO, &buddyRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to map user dto to entity: %w", err)
	}

	return &buddyRequest, nil
}
