package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CacheInit(t *testing.T) {
	cache := NewCache(10)
	cache.Set("hello", "world", 1*time.Second)
	value, ok := cache.Get("hello")
	assert.Equal(t, true, ok, "Value should be available")
	assert.Equal(t, "world", value, "Values should match")
}

func Test_CacheNotAvailable(t *testing.T) {
	cache := NewCache(10)
	value, ok := cache.Get("hello")
	assert.Equal(t, false, ok, "Value should not be available")
	assert.Equal(t, "", value, "Values should not be available")
}

func Test_CacheUpdate(t *testing.T) {
	cache := NewCache(10)

	cache.Set("hello", "world", 1*time.Second)
	value, ok := cache.Get("hello")
	assert.Equal(t, true, ok, "Value should not be available")
	assert.Equal(t, "world", value, "Values should not be available")

	cache.Set("hello", "other world", 1*time.Second)
	value, ok = cache.Get("hello")
	assert.Equal(t, true, ok, "Value should not be available")
	assert.Equal(t, "other world", value, "Values should not be available")
}
