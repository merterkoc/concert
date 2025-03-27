package userservice

import entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"

type userService struct {
	userRepo repository.UserRepository
}

func (e *userService) CreateUser(username string) (entity.User, error) {
	user := entity.User{
		Username: username,
	}
	return e.userRepo.Save(user)
}

func NewUserService(userRepo repository.UserRepository) service.UserService {
	return userService{userRepo: userRepo}
}
