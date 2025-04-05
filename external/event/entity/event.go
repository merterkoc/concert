package entity

type Embedde struct {
	ApiResponse ApiResponse `json:"_embedded"`
}

type ApiResponse struct {
	Events []Event `json:"events" gorm:"-"`
}

type Event struct {
	Name        string       `json:"name" gorm:"-"`
	Type        string       `json:"type" gorm:"-"`
	ID          string       `json:"id" gorm:"-"`
	Locale      string       `json:"locale" gorm:"-"`
	Dates       Dates        `json:"dates" gorm:"-"`
	PriceRanges []PriceRange `json:"priceRanges" gorm:"-"`
	Embedded    Embedded     `json:"_embedded" gorm:"-"`
}

type SalesPeriod struct {
	StartDateTime string `json:"startDateTime"`
	EndDateTime   string `json:"endDateTime"`
	StartTBD      bool   `json:"startTBD"`
	StartTBA      bool   `json:"startTBA"`
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

type Venue struct {
	Name       string  `json:"name"`
	Locale     string  `json:"locale"`
	PostalCode string  `json:"postalCode"`
	City       City    `json:"city"`
	Country    Country `json:"country"`
}
