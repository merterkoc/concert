package dto

import (
	"gilab.com/pragmaticreviews/golang-gin-poc/internal/identity/entity"
)

type GetAllInterestsResponse struct {
	Interests []entity.InterestType `json:"interests"`
}
