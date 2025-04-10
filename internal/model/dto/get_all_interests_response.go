package dto

import "gilab.com/pragmaticreviews/golang-gin-poc/internal/model/entity"

type GetAllInterestsResponse struct {
	Interests []entity.InterestType `json:"interests"`
}
