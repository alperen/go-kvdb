package commands

import (
	"go-kvdb/database"
	"strconv"
)

var (
	errTTLKeyNotDefRes = Response{StatusErr, "Key field is not defined in received arguments", nil}
)
var TTL CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, exists := m["key"]

	if !exists {
		return errTTLKeyNotDefRes, false
	}

	diff, _ := db.GetEntryTTLDuration(key)
	diffSecs := diff.Seconds()
	ttl := strconv.FormatFloat(diffSecs, 'f', 6, 64)

	return Response{StatusOK, "", map[string]string{"ttl": ttl}}, true
}
