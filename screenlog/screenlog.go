package screenlog

import (
	"fmt"
	"go-kvdb/database"
	"os"
	"sync"
	"text/tabwriter"
)

type ScreenLog struct {
	port                   string
	allTimeConnCount       int
	connectedAddrs         []string
	persistToDiskInSeconds int
	dbSizePtr              *func() int
	maxDbSize              int
	dbEntryCountPtr        *func() int

	sync.Mutex
}

func CreateScreenLog(port string, persistToDiskInSeconds int, maxDbSize int, dbSizePtr *func() int, dbEntryCountPtr *func() int) *ScreenLog {
	return &ScreenLog{
		port:                   port,
		allTimeConnCount:       0,
		connectedAddrs:         []string{},
		persistToDiskInSeconds: persistToDiskInSeconds,
		dbSizePtr:              dbSizePtr,
		maxDbSize:              maxDbSize,
		dbEntryCountPtr:        dbEntryCountPtr,
		Mutex:                  sync.Mutex{},
	}
}

func (sclog *ScreenLog) Print() {
	sclog.Lock()
	defer sclog.Unlock()

	dbSize := *sclog.dbSizePtr
	dbEntryCount := *sclog.dbEntryCountPtr

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	fmt.Fprintf(writer, "Port\t%s\n", sclog.port)
	fmt.Fprintf(writer, "Active Connections\t%d\n", len(sclog.connectedAddrs))
	fmt.Fprintf(writer, "Connected Client Addresses\t%v\n", sclog.connectedAddrs)
	fmt.Fprintf(writer, "All Time Connections\t%d\n", sclog.allTimeConnCount)
	fmt.Fprintf(writer, "Persist to Disk In Seconds\t%d\n", sclog.persistToDiskInSeconds)
	fmt.Fprintf(writer, "DB Size\t%d\n", dbSize())

	if sclog.maxDbSize == database.AbsolutDB {
		fmt.Fprintf(writer, "Max DB Size\t[ABSOLUT_DB]\n")
	} else {
		fmt.Fprintf(writer, "Max DB Size\t%d\n", sclog.maxDbSize)
	}

	fmt.Fprintf(writer, "DB Entry Count\t%d\n", dbEntryCount())

	writer.Flush()
}

func (sclog *ScreenLog) AddClientAddr(addr string) {
	sclog.Lock()
	defer sclog.Unlock()

	sclog.connectedAddrs = append(sclog.connectedAddrs, addr)
	sclog.allTimeConnCount++
}

func (sclog *ScreenLog) RemoveClientAddr(addr string) {
	sclog.Lock()
	defer sclog.Unlock()

	for i, v := range sclog.connectedAddrs {
		if v == addr {
			sclog.connectedAddrs = remove(sclog.connectedAddrs, i)
		}
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
