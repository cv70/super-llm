package cache

type ShardCache[V any] struct {
	cache []*Cache[string, V]
}

func NewShardCache[V any](size int) *ShardCache[V] {
	cache := make([]*Cache[string, V], size)
	for i := range size {
		cache[i] = NewCache[string, V]()
	}
	return &ShardCache[V]{cache: cache}
}

func (g *ShardCache[V]) Shard(key string) *Cache[string, V] {
	return g.cache[len(key) & (len(g.cache) - 1)]
}
