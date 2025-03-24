package service

import entity "gilab.com/pragmaticreviews/golang-gin-poc/internal/event/domain"

type EventService interface {
	FindByKeywordOrLocation(keyword string, location string, page int, size int) (entity.Event, error)
}
