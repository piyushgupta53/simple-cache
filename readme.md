# Simple Cache

This project is a simple in-memory cache implementation in Go, featuring both a basic cache and an LRU (Least Recently Used) cache. It supports setting and retrieving items with optional time-to-live (TTL) values and includes automatic cleanup of expired items.

## Features

- **Basic Cache**: Store key-value pairs with optional TTL.
- **LRU Cache**: Evicts the least recently used items when the cache reaches its capacity.
- **Automatic Cleanup**: Periodically removes expired items from the cache.

## Installation

To use this project, you need to have Go installed. Clone the repository and run:

```bash
go mod tidy
```

## Usage

### Basic Cache

The basic cache allows you to set and get items with an optional TTL. If no TTL is provided, a default TTL is used.

#### Example

```go
cache := NewCache(30 * time.Second)
cache.StartCleanup(5 * time.Second)

// Adding items
cache.Set("key1", "value1", 3 * time.Second)
cache.Set("key2", "value2", 0)

// Retrieving items
value, found := cache.Get("key1")
if found {
    fmt.Println("Found key1:", value)
} else {
    fmt.Println("Key1 not found or expired")
}
```

### LRU Cache

```go
lruCache := NewLRUCache(2, cache)

// Adding items
lruCache.Set("key1", "value1", 3 * time.Second)
lruCache.Set("key2", "value2", 0)
lruCache.Set("key3", "value3", 0) // This will evict "key1"

// Retrieving items
value, found := lruCache.Get("key1")
if found {
    fmt.Println("Found key1:", value)
} else {
    fmt.Println("Key1 not found or expired")
}
```
