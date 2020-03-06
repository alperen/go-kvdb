package commands

import "go-kvdb/database"

// Request struct specifies format of each request.
type Request struct {
	Command string            `json:"command"`
	Args    map[string]string `json:"args,omitempty"`
}

// Response struct specifies format of each request. Message or Result file may be nil or empty in a response
type Response struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Result  map[string]string `json:"result,omitempty"`
}

//CommandFunc is an alias type to Command functions
type CommandFunc func(*database.Database, map[string]string) (Response, bool)

var (
	// StatusOK is a placeholder for Response.Status in successful responses.
	StatusOK = "OK"
	// StatusErr is a placeholder for Response.Status in unsuccessful responses.
	StatusErr = "ERROR"
)
