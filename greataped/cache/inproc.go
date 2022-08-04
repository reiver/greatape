package cache

import (
	"contracts"
	"sync"
)

type inproc struct {
	sync.RWMutex
	collection map[string]interface{}
}

func NewInProcCache() contracts.ICache {
	return &inproc{
		collection: make(map[string]interface{}),
	}
}

func (cache *inproc) Put(key string, value interface{}) {
	cache.Lock()
	defer cache.Unlock()

	cache.collection[key] = value
}

func (cache *inproc) Get(key string) interface{} {
	cache.RLock()
	defer cache.RUnlock()

	if item, exists := cache.collection[key]; exists {
		return item
	}

	return nil
}
