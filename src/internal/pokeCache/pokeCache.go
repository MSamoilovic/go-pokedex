package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mutex *sync.RWMutex
	interval time.Duration
}

type cacheEntry struct {
	val []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cache: map[string]cacheEntry{},
		mutex: &sync.RWMutex{},
		interval: interval,
	}

	go cache.ReapLoop()

	return cache
}

func (ch *Cache) Add(key string, val []byte) {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()

	ch.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (ch *Cache)Get(key string) ([]byte, bool) {
	ch.mutex.Lock()
	defer ch.mutex.Unlock()

	entry, ok := ch.cache[key]

	if !ok {
		return []byte{}, false
	}

	return entry.val, true
}

func (ch *Cache) ReapLoop() {
	ticker := time.NewTicker(ch.interval)

	defer ticker.Stop()
	for {
		<-ticker.C

		ch.mutex.Lock()
		for key, entry := range ch.cache {
			if time.Since(entry.createdAt) > ch.interval {
				delete(ch.cache, key)
			}
		}
		ch.mutex.Unlock()
	}
}