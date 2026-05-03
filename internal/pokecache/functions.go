package pokecache
import "time"

var interval time.Duration = 5 * time.Second

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		mapStore: make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.mapStore[key]
	if !exists {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mapStore[key] = cacheEntry{
		createdAt: time.Now(),
		data: data,
	}
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for {
		<-ticker.C
		c.reap()
	}
}

func (c *Cache) reap() {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := time.Now()
	for key, entry := range c.mapStore {
		if now.Sub(entry.createdAt) > c.interval {
			delete(c.mapStore, key)
		}
	}
}