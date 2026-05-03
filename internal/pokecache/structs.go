package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mapStore map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	data []byte
}