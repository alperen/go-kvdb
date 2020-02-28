package commands

import "go-kvdb/database.go"

var Get CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key := m["key"]
	val := db.Get(key)

	return Response{"OK", "", map[string]string{"value": val}}, true
}
