package model

import "time"

type Nasa struct {
	ID             int    `json:"id"`
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	URL            string `json:"url"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type ListNasa struct {
	Data []Nasa
}

type ApodQueryInput struct {
	Date      time.Time `json:"date"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	Count     int       `json:"count"`
	Thumbs    bool      `json:"thumbs"`
}
