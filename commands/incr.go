package commands

import (
	"go-kvdb/database"
	"strconv"
	"strings"
)

var (
	errIncrKeyNotDefRes    = Response{StatusErr, "Key field is not defined in received arguments", nil}
	errIncrKeyNotExistsRes = Response{StatusErr, "There is not any data related with the received key.", nil}
	errIncrParseErrRes     = Response{StatusErr, "Unable to parse the data into number type", nil}
)

/*Incr does increase once the value with related with received key.
 *The holding value could be parsable to float64 or int otherwise throws parse error.
 */
var Incr CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	key, keyExists := m["key"]

	if !keyExists {
		return errIncrKeyNotDefRes, false
	}

	val, valExists := db.Get(key)

	if !valExists {
		return errIncrKeyNotExistsRes, false
	}

	var finalVal string

	if strings.Contains(val, ".") {
		fval, ok := strconv.ParseFloat(val, 64)

		if ok != nil {
			return errIncrParseErrRes, false
		}

		fval++
		finalVal = strconv.FormatFloat(fval, 'f', 6, 64)
	} else {
		ival, ok := strconv.Atoi(val)

		if ok != nil {
			return errIncrParseErrRes, false
		}

		ival++
		finalVal = strconv.Itoa(ival)
	}

	db.Set(key, finalVal)

	return Response{StatusOK, "", map[string]string{"value": finalVal}}, true
}
