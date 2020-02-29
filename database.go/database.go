package database

import (
	"sync"
)

type Database struct {
	entries        map[string]string
	maxSizeInBytes int
	isFull         bool
	entriesWithTTL []string
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

func (db *Database) Get(key string) (string, bool) {
	db.Lock()
	defer db.Unlock()

	value, exists := db.entries[key]

	return value, exists
}

func (db *Database) Delete(key string) bool {
	db.Lock()
	defer db.Unlock()

	delete(db.entries, key)

	return true
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
