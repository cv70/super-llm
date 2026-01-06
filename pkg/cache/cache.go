package cache


import (
	"sync"
)

type Cache[K comparable, V any] struct {
	cache map[K]V
	lock  sync.RWMutex
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		cache: make(map[K]V),
	}
}

func (m *Cache[K, V]) Get(key K) (V, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.cache[key]
	return v, ok
}

func (m *Cache[K, V]) Set(key K, value V) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache[key] = value
}

func (m *Cache[K, V]) Delete(key K) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.cache, key)
}

func (m *Cache[K, V]) Range(f func(key K, value V)) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for key, value := range m.cache {
		f(key, value)
	}
}

func (m *Cache[K, V]) Len() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.cache)
}

func (m *Cache[K, V]) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.cache = make(map[K]V)
}

func (m *Cache[K, V]) Keys() []K {
	m.lock.RLock()
	defer m.lock.RUnlock()
	keys := make([]K, 0, len(m.cache))
	for key := range m.cache {
		keys = append(keys, key)
	}
	return keys
}

func (m *Cache[K, V]) Values() []V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	values := make([]V, 0, len(m.cache))
	for _, value := range m.cache {
		values = append(values, value)
	}
	return values
}

func (m *Cache[K, V]) GetOrSet(key K, defaultValue V) V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.cache[key]
	if ok {
		return v
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	v, ok = m.cache[key]
	if ok {
		return v
	}
	m.cache[key] = defaultValue
	return defaultValue
}

func (m *Cache[K, V]) GetOrSetFunc(key K, defaultValue func() V) V {
	m.lock.RLock()
	defer m.lock.RUnlock()
	v, ok := m.cache[key]
	if ok {
		return v
	}
	m.lock.Lock()
	defer m.lock.Unlock()
	v, ok = m.cache[key]
	if ok {
		return v
	}
	m.cache[key] = defaultValue()
	return defaultValue()
}

func (m *Cache[K, V]) Rdo(f func(map[K]V)) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	f(m.cache)
}

func (m *Cache[K, V]) Wdo(f func(*map[K]V)) {
	m.lock.Lock()
	defer m.lock.Unlock()
	f(&m.cache)
}
