package main

import "time"

type CacheItem struct {
	value      interface{}
	expiration int64
}

func NewCacheItem(value interface{}, ttl time.Duration) *CacheItem {
	var expiration int64

	if ttl > 0 {
		expiration = time.Now().Add(ttl).UnixNano()
	} else {
		expiration = 0
	}

	return &CacheItem{
		value:      value,
		expiration: expiration,
	}
}

func (item *CacheItem) IsExpired() bool {
	if item.expiration == 0 {
		return false
	}

	return time.Now().UnixNano() > item.expiration
}
