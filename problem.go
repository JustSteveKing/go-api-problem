package problem

// APIProblem defines a structure to return if there is a problem with an API request
type APIProblem struct {
	Title  string                  `json:"title,omitempty"`
	Detail string                  `json:"detail,omitempty"`
	Status string                  `json:"status,omitempty"`
	Code   string                  `json:"code,omitempty"`
	Meta   *map[string]interface{} `json:"meta,omitempty"`
}
