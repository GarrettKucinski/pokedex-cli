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
	defer c.mux.Unlock()
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.entries[key] = entry
}

func (c *Cache) Get(key string) (value []byte, found bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if entry, ok := c.entries[key]; ok {
		return entry.val, true
	}

	return
}

func (c *Cache) reap(d time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for key := range c.entries {
		difference := time.Since(c.entries[key].createdAt) > d
		if difference {
			delete(c.entries, key)
		}
	}
}

func (c *Cache) reapLoop(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(d)
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
