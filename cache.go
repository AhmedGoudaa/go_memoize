package go_memoize

import (
	"sync"
	"sync/atomic"
	"time"
)

func zeroValue[T any]() T {
	var zero T
	return zero
}

// cacheGroup manages multiple caches with a shared ticker
type cacheGroup struct {
	now          atomic.Value
	ticker       *time.Ticker
	tickInterval time.Duration
	done         chan struct{}
}

// NewCacheGroup creates a new cache group with shared ticker
func newCacheGroup() *cacheGroup {
	group := &cacheGroup{
		done:         make(chan struct{}),
		tickInterval: time.Millisecond,
	}
	group.now.Store(time.Now().Unix())
	group.startTicker()
	return group
}

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

func (g *cacheGroup) stop() {
	close(g.done)
}

var cacheGroupInstance = newCacheGroup()

type entry[V any] struct {
	value     V
	timeStamp int64
}

type Cache[K comparable, V any] struct {
	Entries    map[K]entry[V]
	ttl        int64
	cacheGroup *cacheGroup
	mu         sync.RWMutex
	zeroVal    V
}

func NewCache[K comparable, V any](ttl int64) *Cache[K, V] {
	return &Cache[K, V]{
		Entries:    make(map[K]entry[V]),
		cacheGroup: cacheGroupInstance,
		ttl:        ttl,
		zeroVal:    zeroValue[V](),
	}
}

func NewCacheSized[K comparable, V any](size int, ttl int64) *Cache[K, V] {
	return &Cache[K, V]{
		Entries:    make(map[K]entry[V], size),
		cacheGroup: cacheGroupInstance,
		ttl:        ttl,
		zeroVal:    zeroValue[V](),
	}
}

func (c *Cache[K, V]) NowUnix() int64 {
	return c.cacheGroup.now.Load().(int64)
}

//func (c *Cache[K, V]) GetOrCompute(key K, computeFn func() V) V {
//	//now := c.NowUnix()
//
//	if val, ok := c.Get(key); ok {
//		return val
//	}
//
//	newVal := computeFn()
//	c.Set(key, newVal)
//	return newVal
//}

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

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	delete(c.Entries, key)
	c.mu.Unlock()
}

func (c *Cache[K, V]) Set(key K, value V) {
	timeStamp := c.NowUnix()
	c.mu.Lock()
	c.Entries[key] = entry[V]{value: value, timeStamp: timeStamp}
	c.mu.Unlock()
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	entry, ok := c.Entries[key]
	c.mu.RUnlock()

	if ok && (c.ttl == 0 || c.NowUnix()-entry.timeStamp < c.ttl) {
		return entry.value, true
	}

	return c.zeroVal, false
}
