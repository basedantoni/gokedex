package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mux  *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		entries: map[string]cacheEntry{},
		interval: interval,
	}

	c.mux.Lock()
	defer c.mux.Unlock()

	c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {

}

func (c *Cache) Get(key string) ([]byte, bool) {
	return []byte{}, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}
}