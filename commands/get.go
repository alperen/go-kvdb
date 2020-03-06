package commands

import "go-kvdb/database"

var (
	errGetKeyNotDefRes = Response{StatusErr, "Key field is not defined in received arguments", nil}
	errValNotFoundRes  = Response{StatusErr, "Key couldn't find in database.", map[string]string{"value": ""}}
)

/*Get does get a value that related with received key.
 *If the key doesn't exist in the database response error message.
 *Reflects the value in result array in the response.
 */
var Get CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]
	val, valExists := db.Get(key)

	if !keyExists {
		return errGetKeyNotDefRes, false
	}

	if !valExists {
		return errValNotFoundRes, false
	}

	return Response{StatusOK, "", map[string]string{"value": val}}, true
}
