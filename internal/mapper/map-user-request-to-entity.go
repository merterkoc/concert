package mapper

import (
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"github.com/mitchellh/mapstructure"
)

func MapUserRequestToEntity(req dto.CreateUserRequest) (entity.User, error) {
	var userEntity entity.User

	err := mapstructure.Decode(req, &userEntity)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to map user request to entity: %w", err)
	}

	return userEntity, nil
}
