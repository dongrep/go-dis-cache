package main

import (
	"container/list"
	"sync"
	"time"
)

type CacheItem struct {
	Value      string
	ExpiryTime time.Time
}

type Cache struct {
	mu       sync.RWMutex
	items    map[string]*list.Element
	eviction list.List
	capacity int
}

type entry struct {
	key   string
	value CacheItem
}

func NewCache(capacity int) *Cache {
	return &Cache{
		items:    make(map[string]*list.Element),
		eviction: *list.New(),
		capacity: capacity,
	}
}

func (c *Cache) Set(key, value string, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.items[key]; found {
		c.eviction.Remove(elem)
		delete(c.items, key)
	}

	if c.eviction.Len() > c.capacity {
		c.evictLRU()
	}

	item := CacheItem{
		Value:      value,
		ExpiryTime: time.Now().Add(ttl),
	}
	elem := c.eviction.PushFront(&entry{key, item})
	c.items[key] = elem
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]
	if !found || time.Now().After(item.Value.(*entry).value.ExpiryTime) {
		if found {
			c.eviction.Remove(item)
			delete(c.items, key)
		}
		return "", false
	}
	c.eviction.MoveToFront(item)
	return item.Value.(*entry).value.Value, true
}

func (c *Cache) startEvictionTicker(d time.Duration) {
	ticker := time.NewTicker(d)

	go func() {
		for range ticker.C {
			c.evictExpiredItems()
		}
	}()
}

func (c *Cache) evictExpiredItems() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, item := range c.items {
		if time.Now().After(item.Value.(*entry).value.ExpiryTime) {
			c.eviction.Remove(item)
			delete(c.items, key)
		}
	}
}

func (c *Cache) evictLRU() {
	elem := c.eviction.Back()
	if elem != nil {
		c.eviction.Remove(elem)
		delete(c.items, elem.Value.(*entry).key)
	}
}
