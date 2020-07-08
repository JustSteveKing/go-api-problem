// Package problem provides a simple struct for creating API Problems for your Go Lang API
package problem

// APIProblem defines a structure to return if there is a problem with an API request
// It takes a Title, Detail, Status and Code all as Strings.
type APIProblem struct {
	// Title is a short, human readable summary of your problem
	Title string `json:"title,omitempty"`

	// Detail is a longer human-readable explaination which explains your problem in more detail
	Detail string `json:"detail,omitempty"`

	// Status is te HTTP status code that is applicable to the problem itself
	Status string `json:"status,omitempty"`

	// Code is the specific application error code - can be used for internal tracing
	Code string `json:"code,omitempty"`

	// Meta is an object containing none specific meta-details about the problem
	Meta *map[string]interface{}
}
