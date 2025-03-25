package dto

type GetEventRequest struct {
	Size     int    `form:"size" binding:"required"`
	Page     int    `form:"page" binding:"required"`
	Keyword  string `form:"keyword"`
	Location string `form:"location"`
}
