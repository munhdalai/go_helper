package cache

import (
	"sync"
	"time"
)

type entry[V any] struct {
	value     V
	expiresAt time.Time
	hasTTL    bool
}

func (e entry[V]) isExpired() bool {
	if !e.hasTTL {
		return false
	}
	return time.Now().After(e.expiresAt)
}

// Cache нь thread-safe generic кэш.
type Cache[K comparable, V any] struct {
	mu         sync.RWMutex
	items      map[K]entry[V]
	defaultTTL time.Duration
}

// New нь шинэ Cache үүсгэнэ. defaultTTL=0 бол хугацаагүй.
func New[K comparable, V any](defaultTTL time.Duration) *Cache[K, V] {
	return &Cache[K, V]{
		items:      make(map[K]entry[V]),
		defaultTTL: defaultTTL,
	}
}

// Set нь утга хадгална (default TTL ашиглана).
func (c *Cache[K, V]) Set(key K, value V) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e := entry[V]{value: value}
	if c.defaultTTL > 0 {
		e.expiresAt = time.Now().Add(c.defaultTTL)
		e.hasTTL = true
	}
	c.items[key] = e
}

// SetWithTTL нь утга хадгалж, тусгай TTL тохируулна.
func (c *Cache[K, V]) SetWithTTL(key K, value V, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e := entry[V]{value: value}
	if ttl > 0 {
		e.expiresAt = time.Now().Add(ttl)
		e.hasTTL = true
	}
	c.items[key] = e
}

// Get нь утга авна. Хугацаа дууссан бол устгаж false буцаана.
func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	e, ok := c.items[key]
	c.mu.RUnlock()
	if !ok {
		var zero V
		return zero, false
	}
	if e.isExpired() {
		c.Delete(key)
		var zero V
		return zero, false
	}
	return e.value, true
}

// Delete нь утга устгана.
func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.items, key)
}

// Clear нь бүх утгыг устгана.
func (c *Cache[K, V]) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = make(map[K]entry[V])
}

// Count нь хадгалагдсан утгуудын тоог буцаана (хугацаа дууссаныг оруулаад).
func (c *Cache[K, V]) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

// Has нь утга байгаа эсэхийг шалгана.
func (c *Cache[K, V]) Has(key K) bool {
	_, ok := c.Get(key)
	return ok
}

// DeleteExpired нь хугацаа дууссан бүх утгыг устгана.
func (c *Cache[K, V]) DeleteExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, e := range c.items {
		if e.isExpired() {
			delete(c.items, k)
		}
	}
}
