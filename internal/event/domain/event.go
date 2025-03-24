package entity

type Embedde struct {
	ApiResponse ApiResponse `json:"_embedded"`
}

type ApiResponse struct {
	Events []Event `json:"events"`
}

type Event struct {
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	ID          string       `json:"id"`
	Locale      string       `json:"locale"`
	Sales       Sales        `json:"sales"`
	Dates       Dates        `json:"dates"`
	PriceRanges []PriceRange `json:"priceRanges"`
	Embedded    Embedded     `json:"_embedded"`
}

type Sales struct {
	Public SalesPeriod `json:"public"`
}

type SalesPeriod struct {
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
	StartTBD      bool   `json:"startTBD"`
	StartTBA      bool   `json:"startTBA"`
}

type Dates struct {
	Start DateStart `json:"start"`
}

type DateStart struct {
	LocalDate      string `json:"localDate"`
	LocalTime      string `json:"localTime"`
	DateTime       string `json:"dateTime"`
	DateTBD        bool   `json:"dateTBD"`
	Approximate    bool   `json:"dateApproximate"`
	NoSpecificTime bool   `json:"noSpecificTime"`
}

type PriceRange struct {
	Min      float64 `json:"min"`
	Max      float64 `json:"max"`
	Currency string  `json:"currency"`
}

type Embedded struct {
	Venues []Venue `json:"venues"`
}

type Venue struct {
	Name       string  `json:"name"`
	Locale     string  `json:"locale"`
	PostalCode string  `json:"postalCode"`
	City       City    `json:"city"`
	Country    Country `json:"country"`
}

type City struct {
	Name string `json:"name"`
}

type Country struct {
	Name string `json:"name"`
	Code string `json:"countryCode"`
}
