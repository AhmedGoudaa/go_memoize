package go_memoize

import (
	"time"
)

// Memoize returns a memoized version of the compute function with a specified TTL.
// V is the type of the value returned by the compute function.
func Memoize[V any](computeFn func() V, ttl time.Duration) func() V {
	cache := NewCacheSized[uint64, V](1, int64(ttl.Seconds()))
	return func() V {
		return cache.GetOrCompute(0, func() V {
			return computeFn()
		})
	}
}

// Memoize1 returns a memoized version of the compute function with a single key and a specified TTL.
// K is the type of the key, and V is the type of the value returned by the compute function.
func Memoize1[K comparable, V any](computeFn func(K) V, ttl time.Duration) func(K) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(k K) V {
		return cache.GetOrCompute(hash1(k), func() V {
			return computeFn(k)
		})
	}
}

// Memoize2 returns a memoized version of the compute function with two keys and a specified TTL.
// K1 and K2 are the types of the keys, and V is the type of the value returned by the compute function.
func Memoize2[K1, K2 comparable, V any](computeFn func(K1, K2) V, ttl time.Duration) func(K1, K2) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(key1 K1, key2 K2) V {
		return cache.GetOrCompute(hash2(key1, key2), func() V {
			return computeFn(key1, key2)
		})
	}
}

// Memoize3 returns a memoized version of the compute function with three keys and a specified TTL.
// K1, K2, and K3 are the types of the keys, and V is the type of the value returned by the compute function.
func Memoize3[K1, K2, K3 comparable, V any](computeFn func(K1, K2, K3) V, ttl time.Duration) func(K1, K2, K3) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(key1 K1, key2 K2, key3 K3) V {
		return cache.GetOrCompute(hash3(key1, key2, key3), func() V {
			return computeFn(key1, key2, key3)
		})
	}
}

// Memoize4 returns a memoized version of the compute function with four keys and a specified TTL.
// K1, K2, K3, and K4 are the types of the keys, and V is the type of the value returned by the compute function.
func Memoize4[K1, K2, K3, K4 comparable, V any](computeFn func(K1, K2, K3, K4) V, ttl time.Duration) func(K1, K2, K3, K4) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(key1 K1, key2 K2, key3 K3, key4 K4) V {
		return cache.GetOrCompute(hash4(key1, key2, key3, key4), func() V {
			return computeFn(key1, key2, key3, key4)
		})
	}
}

// Memoize5 returns a memoized version of the compute function with five keys and a specified TTL.
// K1, K2, K3, K4, and K5 are the types of the keys, and V is the type of the value returned by the compute function.
func Memoize5[K1, K2, K3, K4, K5 comparable, V any](computeFn func(K1, K2, K3, K4, K5) V, ttl time.Duration) func(K1, K2, K3, K4, K5) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(key1 K1, key2 K2, key3 K3, key4 K4, key5 K5) V {
		return cache.GetOrCompute(hash5(key1, key2, key3, key4, key5), func() V {
			return computeFn(key1, key2, key3, key4, key5)
		})
	}
}

// Memoize6 returns a memoized version of the compute function with six keys and a specified TTL.
// K1, K2, K3, K4, K5, and K6 are the types of the keys, and V is the type of the value returned by the compute function.
func Memoize6[K1, K2, K3, K4, K5, K6 comparable, V any](computeFn func(K1, K2, K3, K4, K5, K6) V, ttl time.Duration) func(K1, K2, K3, K4, K5, K6) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(key1 K1, key2 K2, key3 K3, key4 K4, key5 K5, key6 K6) V {
		return cache.GetOrCompute(hash6(key1, key2, key3, key4, key5, key6), func() V {
			return computeFn(key1, key2, key3, key4, key5, key6)
		})
	}
}

// Memoize7 returns a memoized version of the compute function with seven keys and a specified TTL.
// K1, K2, K3, K4, K5, K6, and K7 are the types of the keys, and V is the type of the value returned by the compute function.
func Memoize7[K1, K2, K3, K4, K5, K6, K7 comparable, V any](computeFn func(K1, K2, K3, K4, K5, K6, K7) V, ttl time.Duration) func(K1, K2, K3, K4, K5, K6, K7) V {
	cache := NewCache[uint64, V](int64(ttl.Seconds()))
	return func(key1 K1, key2 K2, key3 K3, key4 K4, key5 K5, key6 K6, key7 K7) V {
		return cache.GetOrCompute(hash7(key1, key2, key3, key4, key5, key6, key7), func() V {
			return computeFn(key1, key2, key3, key4, key5, key6, key7)
		})
	}
}
