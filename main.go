package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"go-kvdb/commands"
	"go-kvdb/database"
	"go-kvdb/screenlog"
)

var eofIdentifier = byte('\n')

var persistToDiskInSeconds int
var maxMemorySizeInBytes int
var defaultTTLInSeconds int
var refreshRateInSeconds int
var port string
var fileStr string
var panics bool
var detachMode bool

var (
	errBadReqRes       = commands.Response{"error", "Received message could not parsed as json.", nil}
	errUndefinedCmdRes = commands.Response{"error", "Received command is not defined.", nil}
)

var db *database.Database
var sclog *screenlog.ScreenLog

var commandsFuncMap = map[string]commands.CommandFunc{
	"GET":    commands.Get,
	"SET":    commands.Set,
	"DELETE": commands.Delete,
	"INCR":   commands.Incr,
	"DECR":   commands.Decr,
	"EXPIRE": commands.Expire,
	"TTL":    commands.TTL,
	"PING":   commands.Ping,
}

func init() {
	flag.IntVar(&persistToDiskInSeconds, "persist-to-disk", 60, "Server stores the database in every given minutes.")
	flag.IntVar(&maxMemorySizeInBytes, "max-mem-size", database.AbsolutDB, "Sets the maximum size of database. Server does not accepts new entries while maximum size is hanging. Default 0 means no limits.")
	flag.IntVar(&refreshRateInSeconds, "refresh-rate", 1, "Sets screen refresh rate in seconds.")
	flag.StringVar(&port, "port", "6379", "Sets serving port. The given port number should be free for communication")
	flag.StringVar(&fileStr, "file", "", "Refers to database's location on the disk. Should be existed file.")
	flag.BoolVar(&panics, "panics", false, "Shows panics.")
	flag.BoolVar(&detachMode, "detach", false, "Prints nothing to screen")

	flag.Parse()
}

func p(err error) {
	if panics {
		panic(err)
	}
}

func cls() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func detectSignalInterrupt(do func()) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		do()
		os.Exit(0)
	}()
}

func printScreen(sclog *screenlog.ScreenLog, done chan bool) {
	for {
		sclog.Print()
		time.Sleep(time.Second * time.Duration(refreshRateInSeconds))
		cls()
	}

	done <- true
}

func main() {
	var dbFilePtr *os.File

	if fileStr == "" {
		tempFile, err := ioutil.TempFile(os.TempDir(), "gokvdb")

		if err != nil {
			log.Fatal("Unable to create db temp file")
			p(err)
		}

		defer tempFile.Close()

		dbFilePtr = tempFile
	} else {
		file, err := os.OpenFile(fileStr, os.O_RDWR, 066)

		if err != nil {
			log.Fatalf("Unable to open db file %s", fileStr)
			p(err)
		}

		defer file.Close()

		dbFilePtr = file
	}

	db = database.CreateDatabase(maxMemorySizeInBytes, persistToDiskInSeconds, dbFilePtr)

	dbSize := db.Size
	dbEntryCount := db.EntryCount
	sclog = screenlog.CreateScreenLog(port, persistToDiskInSeconds, maxMemorySizeInBytes, &dbSize, &dbEntryCount)
	done := make(chan bool)

	listenerAddr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", listenerAddr)

	if err != nil {
		log.Fatalf("Unable to open listener at %s", listenerAddr)
		p(err)
	}

	defer listener.Close()

	detectSignalInterrupt(func() {
		db.WriteDBToFile()
	})

	go waitConnections(listener, done)

	if !detachMode {
		go printScreen(sclog, done)
	}

	go db.TTLWatcher(done)
	go db.PersistToFileWatcher(done)

	<-done
}

func waitConnections(listener net.Listener, done chan bool) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalf("Unable to accept connection. %s", err)
			p(err)
		}

		sclog.AddClientAddr(conn.RemoteAddr().String())

		go listenClient(conn)
	}

	done <- true
}

func listenClient(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadBytes(eofIdentifier)

		if err == io.EOF {
			sclog.RemoveClientAddr(conn.RemoteAddr().String())
			break
		}

		if string(message) == "" {
			continue
		}

		var request commands.Request

		err = json.Unmarshal(message, &request)

		if err != nil {
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
	bytes = append(bytes, eofIdentifier)

	return bytes
}
