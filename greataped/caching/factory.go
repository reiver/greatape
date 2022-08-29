package caching

import . "contracts"

const (
	InProcCache CacheType = iota
)

func CreateCache(componentType CacheType) ICache {
	switch componentType {
	case InProcCache:
		return NewInProcCache()
	default:
		panic("unknown_cache_type")
	}
}
