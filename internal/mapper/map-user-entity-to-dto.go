package mapper

import (
	"fmt"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"github.com/mitchellh/mapstructure"
)

func MapUserEntityToDto(user entity.User) (*dto.UserDto, error) {
	var userDto dto.UserDto

	err := mapstructure.Decode(user, &userDto)
	if err != nil {
		return nil, fmt.Errorf("failed to map user entity to dto: %w", err)
	}

	return &userDto, nil
}
