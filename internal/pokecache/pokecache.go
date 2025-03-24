package pokecache

import (
	"sync"
	"time"
)

func NewCache(ttl time.Duration) *Cache {
	newCache := &Cache{
		data: make(map[string]cacheEntry),
		mutex: sync.Mutex{},
	}
	
	go newCache.reapLoop(ttl)

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	
	
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.data[key]	
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.data {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.data, k)
		}
	}
}