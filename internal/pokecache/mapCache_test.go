package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	cache := NewCache(5 * time.Second)

	key := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	val := []byte(`{"results":[{"name":"canalave-city-area","url":"test-url"}]}`)

	cache.Add(key, val)

	got, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected cache hit, got miss")
	}

	if string(got) != string(val) {
		t.Fatalf("expected %s, got %s", string(val), string(got))
	}
}

func TestCacheMiss(t *testing.T) {
	cache := NewCache(5 * time.Second)

	_, ok := cache.Get("missing-key")
	if ok {
		t.Fatalf("expected cache miss, got hit")
	}
}

func TestCacheReapsOldEntries(t *testing.T) {
	cache := NewCache(50 * time.Millisecond)

	key := "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20"
	val := []byte(`{"results":[{"name":"eterna-forest-area","url":"test-url"}]}`)

	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Fatalf("expected cache hit immediately after add")
	}

	time.Sleep(150 * time.Millisecond)

	_, ok = cache.Get(key)
	if ok {
		t.Fatalf("expected cache entry to be reaped after interval")
	}
}