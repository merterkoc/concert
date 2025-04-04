package mapper

import (
	"fmt"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"

	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/dto"
	"github.com/mitchellh/mapstructure"
)

func MapUserRequestToEntityWithImage(req dto.CreateUserRequest, publicImageUri *string) (entity.User, error) {
	var userEntity entity.User

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           &userEntity,
		ErrorUnused:      false,
	})
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to create decoder: %w", err)
	}

	err = decoder.Decode(req)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to map user request to entity: %w", err)
	}

	if publicImageUri != nil && *publicImageUri != "" {
		userEntity.UserImage = publicImageUri
	} else {
		userEntity.UserImage = nil
	}

	return userEntity, nil
}
