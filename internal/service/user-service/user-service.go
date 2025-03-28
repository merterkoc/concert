package userservice

import (
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/user/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/user/dto"
)

type UserService interface {
	CreateUser(postNewUserRequest dto.PostNewUserRequest) (*entity.User, error)
}
