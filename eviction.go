package main

import (
	"container/list"
	"time"
)

type LRUCache struct {
	capacity int
	cache    *Cache
	order    *list.List
	keys     map[string]*list.Element
}

func NewLRUCache(capacity int, cache *Cache) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    cache,
		order:    list.New(),
		keys:     make(map[string]*list.Element),
	}
}

func (l *LRUCache) Get(key string) (interface{}, bool) {
	value, found := l.cache.Get(key)

	if found {
		l.updateAccessOrder(key)
	}

	return value, found
}

func (l *LRUCache) Set(key string, value interface{}, ttl time.Duration) {
	if _, found := l.cache.Get(key); found {
		l.updateAccessOrder(key)
	} else {
		if l.order.Len() >= l.capacity {
			l.evict()
		}
		element := l.order.PushFront(key)
		l.keys[key] = element
	}

	l.cache.Set(key, value, ttl)
}

func (l *LRUCache) updateAccessOrder(key string) {
	if element, found := l.keys[key]; found {
		l.order.MoveToFront(element)
	}
}

func (l *LRUCache) evict() {
	element := l.order.Back()

	if element != nil {
		key := element.Value.(string)

		delete(l.cache.items, key)
		l.order.Remove(element)
		delete(l.keys, key)
	}
}
