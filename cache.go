package main

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.RWMutex
	items map[string]*CacheItem
	ttl   time.Duration
}

func NewCache(defaultTTL time.Duration) *Cache {
	return &Cache{
		items: make(map[string]*CacheItem),
		ttl:   defaultTTL,
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if ttl == 0 {
		ttl = c.ttl
	}

	item := NewCacheItem(value, ttl)

	c.items[key] = item
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()

	defer c.mu.RUnlock()

	item, ok := c.items[key]

	if !ok || item.IsExpired() {
		if ok {
			delete(c.items, key)
		}
		return nil, false
	}

	return item.value, true
}

func (c *Cache) StartCleanup(interval time.Duration) {
	go func() {

		for {
			time.Sleep(interval)
			c.mu.Lock()
			for key, item := range c.items {
				if item.IsExpired() {
					delete(c.items, key)
				}
			}
			c.mu.Unlock()
		}
	}()
}
