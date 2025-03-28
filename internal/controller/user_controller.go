package controller

import (
	userService "gilab.com/pragmaticreviews/golang-gin-poc/internal/service/user-service"
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/user/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/user/dto"
)

type UserController interface {
	CreateUser(postNewUserRequest dto.PostNewUserRequest) (*entity.User, error)
}

type userController struct {
	userService userService.UserService
}

// CreateUser is a controller method
// that creates a user
// @Summary Create user
// @Description Create user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.PostNewUserRequest true "User"
// @Success 200 {object} entity.User
// @Router /users [post]
func (c userController) CreateUser(postNewUserRequest dto.PostNewUserRequest) (*entity.User, error) {
	return c.userService.CreateUser(postNewUserRequest)
}

func NewUserController(userService userService.UserService) UserController {
	return userController{userService: userService}
}
