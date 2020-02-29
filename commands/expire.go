package commands

import (
	"go-kvdb/database"
	"strconv"
	"time"
)

var (
	errExpireKeyNotDefRes    = Response{StatusErr, "Key field is not defined in received arguments", nil}
	errExpireTTLNotNumberRes = Response{StatusErr, "TTL value should be a number which is integer.", nil}
)

var Expire CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]
	ttl, ttlExists := m["ttl"]

	if !keyExists {
		return errExpireKeyNotDefRes, false
	}

	if !ttlExists {
		return errExpireTTLNotNumberRes, false
	}

	ttlInt, ttlOk := strconv.Atoi(ttl)

	if ttlOk != nil {
		return errExpireTTLNotNumberRes, false
	}

	db.SetTTLValue(key, time.Duration(ttlInt)*time.Second)

	return Response{StatusOK, "", nil}, true
}
