package go_memoize

import (
	"sync"
	"sync/atomic"
	"time"
)

// zeroValue returns the zero value for any type T.
func zeroValue[T any]() T {
	var zero T
	return zero
}

// cacheGroup manages multiple caches with a shared ticker.
type cacheGroup struct {
	now          atomic.Value
	ticker       *time.Ticker
	tickInterval time.Duration
	done         chan struct{}
}

// newCacheGroup creates a new cache group with a shared ticker.
func newCacheGroup() *cacheGroup {
	group := &cacheGroup{
		done:         make(chan struct{}),
		tickInterval: time.Millisecond,
	}
	group.now.Store(time.Now().Unix())
	group.startTicker()
	return group
}

// startTicker starts the ticker for the cache group.
func (g *cacheGroup) startTicker() {
	g.ticker = time.NewTicker(g.tickInterval) // Initialize the ticker
	go func() {
		for {
			select {
			case <-g.ticker.C:
				g.now.Store(time.Now().Unix())
			case <-g.done:
				g.ticker.Stop()
				return
			}
		}
	}()
}

// var cacheGroupInstance is a singleton instance of cacheGroup.
var cacheGroupInstance = newCacheGroup()

// entry represents a cache entry with a value and a timestamp.
type entry[V any] struct {
	value     V
	timeStamp int64
}

// Cache is a generic cache with a time-to-live (TTL) for each entry.
type Cache[K comparable, V any] struct {
	Entries    map[K]entry[V]
	ttl        int64
	cacheGroup *cacheGroup
	mu         sync.RWMutex
	zeroVal    V
}

// NewCache creates a new cache with the specified TTL.
func NewCache[K comparable, V any](ttl int64) *Cache[K, V] {
	return &Cache[K, V]{
		Entries:    make(map[K]entry[V]),
		cacheGroup: cacheGroupInstance,
		ttl:        ttl,
		zeroVal:    zeroValue[V](),
	}
}

// NewCacheSized creates a new cache with the specified size and TTL.
func NewCacheSized[K comparable, V any](size int, ttl int64) *Cache[K, V] {
	return &Cache[K, V]{
		Entries:    make(map[K]entry[V], size),
		cacheGroup: cacheGroupInstance,
		ttl:        ttl,
		zeroVal:    zeroValue[V](),
	}
}

// NowUnix returns the current Unix timestamp from the cache group.
func (c *Cache[K, V]) NowUnix() int64 {
	return c.cacheGroup.now.Load().(int64)
}

// GetOrCompute retrieves the value for the given key or computes it using the provided function if not present or expired.
func (c *Cache[K, V]) GetOrCompute(key K, computeFn func() V) V {
	c.mu.RLock()
	existingEntry, ok := c.Entries[key]
	c.mu.RUnlock()

	now := c.NowUnix()
	if ok && (c.ttl == 0 || now-existingEntry.timeStamp < c.ttl) {
		return existingEntry.value
	}

	c.mu.Lock()
	newVal := computeFn()
	c.Entries[key] = entry[V]{value: newVal, timeStamp: now}
	c.mu.Unlock()
	return newVal
}

// Delete removes the entry for the given key from the cache.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	delete(c.Entries, key)
	c.mu.Unlock()
}

// Set adds or updates the value for the given key in the cache.
func (c *Cache[K, V]) Set(key K, value V) {
	timeStamp := c.NowUnix()
	c.mu.Lock()
	c.Entries[key] = entry[V]{value: value, timeStamp: timeStamp}
	c.mu.Unlock()
}

// Get retrieves the value for the given key from the cache if present and not expired.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	entry, ok := c.Entries[key]
	c.mu.RUnlock()

	if ok && (c.ttl == 0 || c.NowUnix()-entry.timeStamp < c.ttl) {
		return entry.value, true
	}

	return c.zeroVal, false
}
