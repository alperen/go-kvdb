package main

import (
	"flag"
	"log"
	"time"
)

var persistToDiskInSeconds int
var maxMemorySize int
var defaultTTLInSeconds int
var port int
var file string

func init() {
	flag.IntVar(&persistToDiskInSeconds, "persist-to-disk", int(time.Minute), "Server stores the database in every given minutes.")
	flag.IntVar(&maxMemorySize, "max-mem-size", 0, "Sets the maximum size of database. Server does not accepts new entries while maximum size is hanging. Default 0 means no limits.")
	flag.IntVar(&port, "port", 6379, "Sets serving port. The given port number should be free for communication")
	flag.StringVar(&file, "file", "", "Refers to database's location on the disk. Should be existed file.")

	flag.Parse()
}

func main() {

	log.Println(port)

}
