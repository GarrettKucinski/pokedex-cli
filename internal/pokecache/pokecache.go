package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]CacheEntry
	mux     sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = entry
	c.mux.Unlock()
}

func (c *Cache) Get(key string) (value []byte, found bool) {
	c.mux.Lock()
	if entry, ok := c.entries[key]; ok {
		c.mux.Unlock()
		return entry.val, true
	}
	c.mux.Unlock()

	return
}

func (c *Cache) reapLoop(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for range ticker.C {
		for key := range c.entries {
			difference := time.Since(c.entries[key].createdAt) > d
			if difference {
				c.mux.Lock()
				delete(c.entries, key)
				c.mux.Unlock()
			}
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries: make(map[string]CacheEntry),
		mux:     sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
