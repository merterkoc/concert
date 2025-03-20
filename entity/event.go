package entity

type Event struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Date     string `json:"date"`
}
