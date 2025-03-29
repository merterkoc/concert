package eventservice

import entity "gilab.com/pragmaticreviews/golang-gin-poc/external/event/domain"

type EventService interface {
	FindById(id string) (entity.EventDetail, error)
	FindByKeywordOrLocation(keyword string, location string, page int, size int) ([]entity.Event, error)
}
