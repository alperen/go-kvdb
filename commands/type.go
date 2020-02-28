package commands

import "go-kvdb/database.go"

type Request struct {
	Command string            `json:"command"`
	Args    map[string]string `json:"args,omitempty"`
}

type Response struct {
	Status  string            `json:"status"`
	Message string            `json:"message,omitempty"`
	Result  map[string]string `json:"result,omitempty"`
}

type CommandFunc func(*database.Database, map[string]string) (Response, bool)
