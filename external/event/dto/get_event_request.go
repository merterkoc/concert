package dto

type GetEventRequest struct {
	Size     int    `form:"size" binding:"required"`
	Page     int    `form:"page" binding:"omitempty,min=0"`
	Keyword  string `form:"keyword"`
	Location string `form:"location"`
}
