package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"go-kvdb/commands"

	"go-kvdb/database"
)

var jsonEOF = '}'

var persistToDiskInSeconds int
var maxMemorySize float64
var defaultTTLInSeconds int
var port string
var file string
var panics bool

var (
	errBadReqRes       = commands.Response{"error", "Received message could not parsed as json.", nil}
	errUndefinedCmdRes = commands.Response{"error", "Received command is not defined.", nil}
)

var db *database.Database

var commandsFuncMap = map[string]commands.CommandFunc{
	"GET":    commands.Get,
	"SET":    commands.Set,
	"DELETE": commands.Delete,
	"INCR":   commands.Incr,
	"EXPIRE": commands.Expire,
	"TTL":    commands.TTL,
	"PING":   commands.Ping,
}

func init() {
	flag.IntVar(&persistToDiskInSeconds, "persist-to-disk", int(time.Minute), "Server stores the database in every given minutes.")
	flag.Float64Var(&maxMemorySize, "max-mem-size", 0, "Sets the maximum size of database. Server does not accepts new entries while maximum size is hanging. Default 0 means no limits.")
	flag.StringVar(&port, "port", "6379", "Sets serving port. The given port number should be free for communication")
	flag.StringVar(&file, "file", "", "Refers to database's location on the disk. Should be existed file.")
	flag.BoolVar(&panics, "panics", true, "Shows panics.")

	flag.Parse()
}

func p(err error) {
	if panics {
		panic(err)
	}
}

func main() {

	db = database.CreateDatabase(100)

	listenerAddr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", listenerAddr)

	if err != nil {
		log.Fatalf("Unable to open listener at %s", listenerAddr)
		p(err)
	}

	defer listener.Close()

	waitConnections(listener)
}

func waitConnections(listener net.Listener) {
	log.Println("Waits connections")
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("Unable to accept connection. %s", err)
			p(err)
		}

		log.Printf("connection %s", conn.RemoteAddr())

		go listenClient(conn)
	}
}

func listenClient(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadBytes('\n')

		if err == io.EOF {
			log.Print("Connection lost")
			break
		}

		if string(message) == "" {
			continue
		}

		var request commands.Request

		err = json.Unmarshal(message, &request)

		if err != nil {
			log.Println(err, string(message))
			conn.Write(bytedRes(errBadReqRes))
			continue
		}

		command, exists := commandsFuncMap[request.Command]

		if !exists {
			conn.Write(bytedRes(errUndefinedCmdRes))
			continue
		}

		response, _ := command(db, request.Args)

		conn.Write(bytedRes(response))
	}
}

func bytedRes(res commands.Response) []byte {
	bytes, _ := json.Marshal(res)
	bytes = append(bytes, '\n')

	return bytes
}
