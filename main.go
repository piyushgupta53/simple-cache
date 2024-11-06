package main

import (
	"fmt"
	"time"
)

func main() {
	cache := NewCache(30 * time.Second)
	cache.StartCleanup(5 * time.Second)

	// Adding items
	cache.Set("key1", "value1", 3*time.Second)
	cache.Set("key2", "value2", 0)

	// Retrieving items
	value, found := cache.Get("key1")
	if found {
		fmt.Println("Found key1:", value)
	} else {
		fmt.Println("Key1 not found or expired")
	}

	// Wait to see the cleanup in action
	time.Sleep(5 * time.Second)

	value, found = cache.Get("key1")
	if found {
		fmt.Println("Found key1:", value)
	} else {
		fmt.Println("Key1 not found or expired after cleanup")
	}
}
