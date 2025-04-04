package dto

import "errors"

type GetEventRequest struct {
	Size     int    `form:"size"`
	Page     int    `form:"page"`
	Keyword  string `form:"keyword"`
	Location string `form:"location"`
}

func (r *GetEventRequest) Validate() error {
	if r.Keyword == "" {
		if r.Size <= 0 {
			return errors.New("size is required when keyword is not provided")
		}
		if r.Page < 0 {
			return errors.New("page must be 0 or greater")
		}
	}
	return nil
}
