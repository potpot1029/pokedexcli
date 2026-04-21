package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries  map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cc := Cache{
		entries:  map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go cc.reapLoop()

	return cc
}

func (cc Cache) Add(key string, val []byte) {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	cc.entries[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (cc Cache) Get(key string) ([]byte, bool) {
	cc.mu.Lock()
	defer cc.mu.Unlock()

	entry, ok := cc.entries[key]
	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (cc Cache) reapLoop() {
	ticker := time.NewTicker(cc.interval)
	for range ticker.C {
		cc.mu.Lock()

		for key, entry := range cc.entries {
			// check if expired
			if entry.createdAt.Add(cc.interval).Before(time.Now()) {
				delete(cc.entries, key)
			}
		}

		cc.mu.Unlock()
	}
}
