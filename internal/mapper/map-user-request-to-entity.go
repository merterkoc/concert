package mapper

import (
	"fmt"

	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/user/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/user/dto"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashed), nil
}

func MapUserRequestToEntity(req dto.PostNewUserRequest) (entity.User, error) {
	var userEntity entity.User

	err := mapstructure.Decode(req, &userEntity)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to map user request to entity: %w", err)
	}

	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	userEntity.PasswordHash = hashedPassword

	return userEntity, nil
}
