package model

import "time"

// Nasa struct represents mandatory form information for creation
// swagger:model
type Nasa struct {
	// The id of a nasa model
	// example: 1
	// required: false
	ID int `json:"id"`
	// The copyright of a nasa model
	// example: some name
	// required: true
	Copyright string `json:"copyright"`
	// The date of a nasa model
	// example: 2022-08-18
	// required: true
	Date string `json:"date"`
	// The explanation of a nasa model
	// example: some text
	// required: true
	Explanation string `json:"explanation"`
	// The hdurl of a nasa model
	// example: https://apod.nasa.gov/apod/image/2208/perseids2022jcc2k.jpg
	// required: true
	Hdurl string `json:"hdurl"`
	// The media_type of a nasa model
	// example: image
	// required: true
	MediaType string `json:"media_type"`
	// The service_version of a nasa model
	// example: v1
	// required: true
	ServiceVersion string `json:"service_version"`
	// The title of a nasa model
	// example: some text
	// required: true
	Title string `json:"title"`
	// The url of a nasa model
	// example: https://apod.nasa.gov/apod/image/2208/perseids2022jcc2k800.jpg
	// required: true
	URL string `json:"url"`
}

// ApodQueryInput struct represents mandatory model filter information for request
// swagger:model
type ApodQueryInput struct {
	// The date of a nasa filter model
	// example: 2022-08-18
	// required: true
	Date time.Time `json:"date"`
	// The start_date of a nasa filter model
	// example: 2022-08-18
	// required: true
	StartDate time.Time `json:"start_date"`
	// The end_date of a nasa filter model
	// example: 2022-08-18
	// required: true
	EndDate time.Time `json:"end_date"`
	// The count of a nasa filter model
	// example: 8
	// required: true
	Count int `json:"count"`
	// The thumbs of a nasa filter model
	// example: true
	// required: true
	Thumbs bool `json:"thumbs"`
}

// ListNasa struct represents a list of data
// swagger:model
type ListNasa struct {
	Data []Nasa
}

type ErrorResponse struct {
	Message string `json:"message"`
}
