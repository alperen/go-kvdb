package commands

import "go-kvdb/database"

var (
	errDelKeyNotDefRes = Response{StatusErr, "Key field is not defined in received arguments", nil}
)

/*Delete does delete the value that related with received key.
 *Always return OK response.
 */
var Delete CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]

	if !keyExists {
		return errDelKeyNotDefRes, false
	}

	ok := db.Delete(key)

	return Response{StatusOK, "", nil}, ok
}
