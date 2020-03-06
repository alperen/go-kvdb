package commands

import (
	"go-kvdb/database"
)

/*Ping is a just basic implementation that to test connection.
 *Reflects PONG data.
 */
var Ping CommandFunc = func(db *database.Database, m map[string]string) (Response, bool) {
	return Response{StatusOK, "PONG", nil}, true
}
