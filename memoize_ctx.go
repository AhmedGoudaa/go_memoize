package go_memoize

import (
	"context"
	"time"
)

// MemoizeCtx returns a memoized version of the compute function with a specified TTL.
func MemoizeCtx[V any](computeFn func(context.Context) V, ttl time.Duration) func(context.Context) V {
	cache := NewCacheSized[uint64, V](1, int64(ttl.Seconds()))
	return func(ctx context.Context) V {
		return cache.GetOrCompute(0, func() V {
			return computeFn(ctx)
		})
	}
}

// MemoizeCtx1 returns a memoized version of the compute function with a single key and a specified TTL.
func MemoizeCtx1[K comparable, V any](computeFn func(context.Context, K) V, ttl time.Duration) func(context.Context, K) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, k K) V {
		return cache.GetOrCompute(hash1(k), func() V {
			return computeFn(ctx, k)
		})
	}
}

// MemoizeCtx2 returns a memoized version of the compute function with two keys and a specified TTL.
func MemoizeCtx2[K1, K2 comparable, V any](computeFn func(context.Context, K1, K2) V, ttl time.Duration) func(context.Context, K1, K2) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, key1 K1, key2 K2) V {
		return cache.GetOrCompute(hash2(key1, key2), func() V {
			return computeFn(ctx, key1, key2)
		})
	}
}

// MemoizeCtx3 returns a memoized version of the compute function with three keys and a specified TTL.
func MemoizeCtx3[K1, K2, K3 comparable, V any](computeFn func(context.Context, K1, K2, K3) V, ttl time.Duration) func(context.Context, K1, K2, K3) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, key1 K1, key2 K2, key3 K3) V {
		return cache.GetOrCompute(hash3(key1, key2, key3), func() V {
			return computeFn(ctx, key1, key2, key3)
		})
	}
}

// MemoizeCtx4 returns a memoized version of the compute function with four keys and a specified TTL.
func MemoizeCtx4[K1, K2, K3, K4 comparable, V any](computeFn func(context.Context, K1, K2, K3, K4) V, ttl time.Duration) func(context.Context, K1, K2, K3, K4) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, key1 K1, key2 K2, key3 K3, key4 K4) V {
		return cache.GetOrCompute(hash4(key1, key2, key3, key4), func() V {
			return computeFn(ctx, key1, key2, key3, key4)
		})
	}
}

// MemoizeCtx5 returns a memoized version of the compute function with five keys and a specified TTL.
func MemoizeCtx5[K1, K2, K3, K4, K5 comparable, V any](computeFn func(context.Context, K1, K2, K3, K4, K5) V, ttl time.Duration) func(context.Context, K1, K2, K3, K4, K5) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, key1 K1, key2 K2, key3 K3, key4 K4, key5 K5) V {
		return cache.GetOrCompute(hash5(key1, key2, key3, key4, key5), func() V {
			return computeFn(ctx, key1, key2, key3, key4, key5)
		})
	}
}

// MemoizeCtx6 returns a memoized version of the compute function with six keys and a specified TTL.
func MemoizeCtx6[K1, K2, K3, K4, K5, K6 comparable, V any](computeFn func(context.Context, K1, K2, K3, K4, K5, K6) V, ttl time.Duration) func(context.Context, K1, K2, K3, K4, K5, K6) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, key1 K1, key2 K2, key3 K3, key4 K4, key5 K5, key6 K6) V {
		return cache.GetOrCompute(hash6(key1, key2, key3, key4, key5, key6), func() V {
			return computeFn(ctx, key1, key2, key3, key4, key5, key6)
		})
	}
}

// MemoizeCtx7 returns a memoized version of the compute function with seven keys and a specified TTL.
func MemoizeCtx7[K1, K2, K3, K4, K5, K6, K7 comparable, V any](computeFn func(context.Context, K1, K2, K3, K4, K5, K6, K7) V, ttl time.Duration) func(context.Context, K1, K2, K3, K4, K5, K6, K7) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(ctx context.Context, key1 K1, key2 K2, key3 K3, key4 K4, key5 K5, key6 K6, key7 K7) V {
		return cache.GetOrCompute(hash7(key1, key2, key3, key4, key5, key6, key7), func() V {
			return computeFn(ctx, key1, key2, key3, key4, key5, key6, key7)
		})
	}
}
