package errors

// Error describes an HTTP error to be sent to a client.
type Error struct {
	// Status is an HTTP status code (ex: 200).
	Status int `json:"status"`
	// Code is an internal error code.
	Code string `json:"code"`
	// Message is a human readable message.
	Message string `json:"message"`
}
