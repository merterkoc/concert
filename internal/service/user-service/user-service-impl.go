package userservice

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/mapper"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/repository"
	entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/user/domain"
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/user/dto"
)

type userService struct {
	userRepo repository.UserRepository
}

func (s *userService) CreateUser(request dto.PostNewUserRequest) (*entity.User, error) {
	entity, err := mapper.MapUserRequestToEntity(request)
	if err != nil {
		return nil, err
	}
	res, err := s.userRepo.SaveUser(&entity)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func NewUserService(
	userRepo *repository.UserRepository,
) UserService {
	return &userService{
		userRepo: *userRepo,
	}
}
