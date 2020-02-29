package commands

import "go-kvdb/database"

var TTL CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	return Response{}, true
}
