package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu         sync.Mutex
	interval   time.Duration
	cacheEntry map[string]CacheEntry
}
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		mu:         sync.Mutex{},
		interval:   interval,
		cacheEntry: make(map[string]CacheEntry),
	}
	go cache.reapLoop()
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntry[key] = CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	i, ok := c.cacheEntry[key]
	c.mu.Unlock()
	if ok {
		return i.val, true
	}
	return nil, false
}

func (c *Cache) remove(key string) {
	c.mu.Lock()
	delete(c.cacheEntry, key)
	c.mu.Unlock()
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	done := make(chan bool)
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			// reap old entries
			for name, entry := range c.cacheEntry {
				difference := time.Since(entry.createdAt)
				if difference >= c.interval {
					fmt.Printf("Entry for %v expired at %v\n", name, t)
					c.remove(name)
					fmt.Printf("Removed %v\n", name)
				}
			}
		}
	}
}
