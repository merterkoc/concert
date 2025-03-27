package userservice

import entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"

type EventService interface {
	CreateUser(username string) (entity.User, error)
}
