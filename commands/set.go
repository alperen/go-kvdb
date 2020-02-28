package commands

import "go-kvdb/database.go"

var (
	errSetKeyNotDefRes = Response{"error", "Key field is not defined in received arguments", nil}
	errSetValNotDefRes = Response{"error", "Value field is not defined in received arguments", nil}
)

var Set CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]
	val, valExists := m["value"]

	if !keyExists {
		return errSetKeyNotDefRes, false
	}

	if !valExists {
		return errSetValNotDefRes, false
	}

	_ = db.Set(key, val)

	return Response{"OK", "", map[string]string{"value": val}}, true
}
