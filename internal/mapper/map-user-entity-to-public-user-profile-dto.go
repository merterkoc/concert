package mapper

import (
	"fmt"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/dto"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"
	"github.com/mitchellh/mapstructure"
)

func MapUserEntityToPublicUserProfileDto(user entity.User) (*dto.PublicUserProfileDTO, error) {
	var publicUserProfileResponse dto.PublicUserProfileDTO

	err := mapstructure.Decode(user, &publicUserProfileResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to map user entity to dto: %w", err)
	}

	return &publicUserProfileResponse, nil
}
