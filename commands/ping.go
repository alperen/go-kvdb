package commands

var Ping CommandFunc = func(m map[string]string) (Response, bool) {
	return Response{"OK", "PONG", nil}, true
}
