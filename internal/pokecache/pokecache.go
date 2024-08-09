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

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	done := make(chan bool)
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	done <- true
	// }()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
			fmt.Printf("Current cache: %+v\n", c)
		}
	}
}
