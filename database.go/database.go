package database

import (
	"sync"
)

type Database struct {
	entries        map[string]string
	maxSizeInBytes int
	isFull         bool
	sync.Mutex
}

func CreateDatabase(maxSize int) *Database {
	return &Database{
		entries:        make(map[string]string),
		maxSizeInBytes: maxSize,
		isFull:         false,
		Mutex:          sync.Mutex{},
	}
}

func (db *Database) Set(key, value string) bool {
	db.Lock()
	defer db.Unlock()

	db.entries[key] = value

	return true
}

func (db *Database) Get(key string) string {
	db.Lock()
	defer db.Unlock()

	value, _ := db.entries[key]

	return value
}

func (db *Database) Size() int {
	db.Lock()
	defer db.Unlock()
	count := 0

	for k, v := range db.entries {
		count += len(k) + len(v)
	}

	return count
}
