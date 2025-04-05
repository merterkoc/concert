package entity

// Event struct
type EventDetail struct {
	Name            string           `json:"name"`
	Type            string           `json:"type"`
	ID              string           `json:"id"`
	Test            bool             `json:"test"`
	URL             string           `json:"url"`
	Locale          string           `json:"locale"`
	Images          []Image          `json:"images"`
	Dates           Dates            `json:"dates" gorm:"-"`
	Classifications []Classification `json:"classifications" gorm:"-"`
	Links           Links            `json:"_links" gorm:"-"`
	Embedded        Embedded         `json:"_embedded" gorm:"-"`
}

// Image struct
type Image struct {
	Ratio    string `json:"ratio"`
	URL      string `json:"url"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Fallback bool   `json:"fallback"`
}

// PublicSale struct
type PublicSale struct {
	StartDateTime string `json:"startDateTime"`
	StartTBD      bool   `json:"startTBD"`
	StartTBA      bool   `json:"startTBA"`
	EndDateTime   string `json:"endDateTime"`
}

// Dates struct
type Dates struct {
	Start            StartDate `json:"start"`
	Status           Status    `json:"status"`
	Timezone         string    `json:"timezone"`
	SpanMultipleDays bool      `json:"spanMultipleDays"`
}

// StartDate struct
type StartDate struct {
	LocalDate      string `json:"localDate"`
	LocalTime      string `json:"localTime"`
	DateTime       string `json:"dateTime"`
	DateTBD        bool   `json:"dateTBD"`
	DateTBA        bool   `json:"dateTBA"`
	TimeTBA        bool   `json:"timeTBA"`
	NoSpecificTime bool   `json:"noSpecificTime"`
}

// Status struct
type Status struct {
	Code string `json:"code"`
}

// Classification struct
type Classification struct {
	Primary  bool     `json:"primary"`
	Segment  Segment  `json:"segment"`
	Genre    Genre    `json:"genre"`
	SubGenre SubGenre `json:"subGenre"`
	Type     Type     `json:"type"`
	SubType  SubType  `json:"subType"`
	Family   bool     `json:"family"`
}

// Segment struct
type Segment struct {
	ID string `json:"id"`
}

// Genre struct
type Genre struct {
	ID string `json:"id"`
}

// SubGenre struct
type SubGenre struct {
	ID string `json:"id"`
}

// Type struct
type Type struct {
	ID string `json:"id"`
}

// SubType struct
type SubType struct {
	ID string `json:"id"`
}

// Links struct
type Links struct {
	Self        Link   `json:"self"`
	Attractions []Link `json:"attractions"`
	Venues      []Link `json:"venues"`
}

// Link struct
type Link struct {
	Href string `json:"href"`
}

// Embedded struct
type Embedded struct {
	Venues      []Venue      `json:"venues"`
	Attractions []Attraction `json:"attractions"`
}

// Venue struct
type VenueDetail struct {
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	ID             string         `json:"id"`
	Test           bool           `json:"test"`
	URL            string         `json:"url"`
	Locale         string         `json:"locale"`
	PostalCode     string         `json:"postalCode"`
	Timezone       string         `json:"timezone"`
	City           City           `json:"city"`
	Country        Country        `json:"country"`
	Address        Address        `json:"address"`
	Location       Location       `json:"location"`
	Markets        []Market       `json:"markets"`
	Dmas           []Dma          `json:"dmas"`
	UpcomingEvents UpcomingEvents `json:"upcomingEvents"`
	Links          Links          `json:"_links"`
}

// City struct
type City struct {
	Name string `json:"name"`
}

// Country struct
type Country struct {
	Name        string `json:"name"`
	CountryCode string `json:"countryCode"`
}

// Address struct
type Address struct {
	Line1 string `json:"line1"`
}

// Location struct
type Location struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

// Market struct
type Market struct {
	ID string `json:"id"`
}

// Dma struct
type Dma struct {
	ID string `json:"id"`
}

// UpcomingEvents struct
type UpcomingEvents struct {
	WtsTr int `json:"wts-tr"`
}

// Attraction struct
type Attraction struct {
	Type            string           `json:"type"`
	ID              string           `json:"id"`
	Test            bool             `json:"test"`
	URL             string           `json:"url"`
	ExternalLinks   ExternalLinks    `json:"externalLinks"`
	Images          []Image          `json:"images"`
	Classifications []Classification `json:"classifications"`
	UpcomingEvents  UpcomingEvents   `json:"upcomingEvents"`
	Links           Links            `json:"_links"`
}

// ExternalLinks struct
type ExternalLinks struct {
	Facebook    []Link `json:"facebook"`
	Wiki        []Link `json:"wiki"`
	Musicbrainz []Link `json:"musicbrainz"`
	Homepage    []Link `json:"homepage"`
}
