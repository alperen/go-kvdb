package commands

import (
	"go-kvdb/database"
	"strings"
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
	diffStr := strings.TrimSuffix(diff.String(), "s")

	return Response{StatusOK, "", map[string]string{"ttl": diffStr}}, true
}
