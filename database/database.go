package database

import (
	"sync"
	"time"
)

var AbsolutDB = 0

var isAbsolutDB = func(maxSize int) bool { return maxSize == AbsolutDB }

type Database struct {
	entries        map[string]string
	maxSizeInBytes int
	isFull         bool
	entriesWithTTL map[string]time.Time
	sync.Mutex
}

func CreateDatabase(maxSize int) *Database {
	return &Database{
		entries:        make(map[string]string),
		maxSizeInBytes: maxSize,
		entriesWithTTL: make(map[string]time.Time),
		Mutex:          sync.Mutex{},
	}
}

func (db *Database) Set(key, value string) bool {
	if !isAbsolutDB(db.maxSizeInBytes) {
		size := db.Size()
		maxSize := db.maxSizeInBytes
		candidateEntrySize := memoryCalcForEntry(key, value)

		if size+candidateEntrySize > maxSize {
			return false
		}
	}

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

	_, hasTTL := db.entriesWithTTL[key]

	if hasTTL {
		delete(db.entriesWithTTL, key)
	}

	return true
}

func (db *Database) Size() int {
	db.Lock()
	defer db.Unlock()
	count := 0

	for k, v := range db.entries {
		count += memoryCalcForEntry(k, v)
	}

	return count
}

func (db *Database) EntryCount() int {
	db.Lock()
	defer db.Unlock()

	return len(db.entries)
}

func (db *Database) SetTTLValue(key string, duration time.Duration) {
	db.Lock()
	defer db.Unlock()

	now := time.Now()
	expiresAt := now.Add(duration)

	db.entriesWithTTL[key] = expiresAt
}

func (db *Database) GetEntryTTLDuration(key string) (time.Duration, bool) {
	db.Lock()
	defer db.Unlock()

	now := time.Now()
	expiresAt, exists := db.entriesWithTTL[key]

	if !exists {
		return 0, false
	}

	diff := expiresAt.Sub(now)

	return diff, true
}

func (db *Database) TTLWatcher(done chan bool) {
	timer := time.Tick(1 * time.Second)
	for range timer {
		now := time.Now()

		for key, expire := range db.entriesWithTTL {
			diff := expire.Sub(now)

			if diff < 0 {
				db.Delete(key)
			}
		}
	}

	done <- true
}

func memoryCalcForEntry(key, val string) int {
	return len(key) + len(val)
}
