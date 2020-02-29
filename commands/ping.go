package commands

import (
	"go-kvdb/database"
)

var Ping CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	return Response{StatusOK, "PONG", nil}, true
}
