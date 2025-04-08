package dto

type PatchUserInterestsRequest struct {
	Operation  string `json:"operation" binding:"required,oneof=add remove" swagger:"enum(add, remove)"`
	InterestID int    `json:"interest_id" binding:"required,gt=0"`
}
