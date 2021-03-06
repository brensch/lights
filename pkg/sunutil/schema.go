package sunutil

import (
	"time"
)

const (
	endPoint = "https://api.sunrise-sunset.org/json"

	latitude  = "37.3996161838102"
	longitude = "-122.10145712094311"
	timeZone  = "America/Los_Angeles"
)

type APIResponse struct {
	Results *APIResults `json:"results"`
	Status  string      `json:"status"`
}

type JSONTime time.Time

type APIResults struct {
	Sunrise                   time.Time `json:"sunrise"`
	Sunset                    time.Time `json:"sunset"`
	SolarNoon                 time.Time `json:"solar_noon"`
	DayLength                 int       `json:"day_length"`
	CivilTwilightBegin        time.Time `json:"civil_twilight_begin"`
	CivilTwilightEnd          time.Time `json:"civil_twilight_end"`
	NauticalTwilightBegin     time.Time `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       time.Time `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin time.Time `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   time.Time `json:"astronomical_twilight_end"`
}
