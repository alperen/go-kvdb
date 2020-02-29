package commands

import "go-kvdb/database.go"

var (
	errSetKeyNotDefRes = Response{StatusErr, "Key field is not defined in received arguments", nil}
	errSetValNotDefRes = Response{StatusErr, "Value field is not defined in received arguments", nil}
	errSetFailedRes    = Response{StatusErr, "Unable to add key to database. Database may be full.", nil}
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

	ok := db.Set(key, val)

	if ok {
		return Response{StatusOK, "", map[string]string{"value": val}}, true
	}

	return errSetFailedRes, false
}
