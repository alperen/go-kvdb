package commands

import (
	"go-kvdb/database"
	"strconv"
	"time"
)

var (
	errSetKeyNotDefRes = Response{StatusErr, "Key field is not defined in received arguments", nil}
	errSetValNotDefRes = Response{StatusErr, "Value field is not defined in received arguments", nil}
	errSetFailedRes    = Response{StatusErr, "Unable to add key to database. Database may be full.", nil}
	errTTLNotNumberRes = Response{StatusErr, "TTL value should be a number which is integer.", nil}
)

/*Set does create a key-value pair in the database.
 *If the key exists already in the database, the newest value will rewrite on key.
 */
var Set CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]
	val, valExists := m["value"]
	ttl, ttlExists := m["ttl"]

	if !keyExists {
		return errSetKeyNotDefRes, false
	}

	if !valExists {
		return errSetValNotDefRes, false
	}

	ok := db.Set(key, val)

	if ok && ttlExists {
		ttl, err := strconv.Atoi(ttl)

		if err != nil {
			return errTTLNotNumberRes, false
		}

		db.SetTTLValue(key, time.Duration(ttl)*time.Second)
	}

	if ok {
		payload := map[string]string{"value": val}

		if ttlExists {
			payload["ttl"] = ttl
		}

		return Response{StatusOK, "", payload}, true
	}

	return errSetFailedRes, false
}
