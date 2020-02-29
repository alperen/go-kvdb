package sclg

import (
	"log"
	"sync"
)

type ScLg struct {
	port                   string
	activeConnCount        int
	allTimeConnCount       int
	connectedAddrs         []string
	persistToDiskInSeconds int
	dbSize                 float64
	maxDbSize			   float64
	dbEntryCount           int

	sync.Mutex
}

func CreateScreenLog(port string, persistToDiskInSeconds int, maxDbSize float64) ScLg {
	return ScLg{
		port:                   port,
		activeConnCount:        0,
		allTimeConnCount:       0,
		connectedAddrs:         []string{},
		persistToDiskInSeconds: persistToDiskInSeconds,
		dbSize:                 0,
		maxDbSize:              maxDbSize,
		dbEntryCount:           0,
		Mutex:                  sync.Mutex{},
	}
}

func (sclog *ScLg) Print() {
	log.Println("Hello")
}

