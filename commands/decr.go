package commands

import (
	"go-kvdb/database"
	"strconv"
	"strings"
)

var (
	errDecrKeyNotDefRes    = Response{StatusErr, "Key field is not defined in received arguments", nil}
	errDecrKeyNotExistsRes = Response{StatusErr, "There is not any data related with the received key.", nil}
	errDecrParseErrRes     = Response{StatusErr, "Unable to parse the data into number type", nil}
)
var Decr CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]

	if !keyExists {
		return errDecrKeyNotDefRes, false
	}

	val, valExists := db.Get(key)

	if !valExists {
		return errDecrKeyNotExistsRes, false
	}

	var finalVal string

	if strings.Contains(val, ".") {
		fval, ok := strconv.ParseFloat(val, 64)

		if ok != nil {
			return errDecrParseErrRes, false
		}

		fval--
		finalVal = strconv.FormatFloat(fval, 'f', 6, 64)
	} else {
		ival, ok := strconv.Atoi(val)

		if ok != nil {
			return errDecrParseErrRes, false
		}

		ival--
		finalVal = strconv.Itoa(ival)
	}

	db.Set(key, finalVal)

	return Response{StatusOK, "", map[string]string{"value": finalVal}}, true
}
