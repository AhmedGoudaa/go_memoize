package go_memoize

import (
	"fmt"
	"math"
)

// we are using FNV-1a hash algorithm to hash the key
// FNV (Fowler-Noll-Vo) hash algorithm has two main components:
// 1. offset basis
// The offset basis is the starting value of the hash.
// 2. prime multiplier
//The prime multiplier is a large prime number that is used to hash the input data.

// how it works:
// 1. The offset basis is multiplied by the prime multiplier.
// 2. The result is XORed with the first byte of the input data.
// 3. The result is then multiplied by the prime multiplier.
// 4. This process is repeated for each byte of the input data.
// 5. The final result is the hash value.

const (
	// FNV-1a hash constants for 32-bit and 64-bit architectures
	offset64 = uint64(14695981039346656037)
	prime64  = uint64(1099511628211)

	// for hashing boolean values
	trueHash  = offset64 ^ 1*prime64
	falseHash = offset64 ^ 0*prime64
)

func hash1[A comparable](key A) uint64 {
	return hash(offset64, key)
}

func hash2[A, B comparable](key1 A, key2 B) uint64 {
	return hash(hash(offset64, key1), key2)
}

func hash3[A, B, C comparable](key1 A, key2 B, key3 C) uint64 {
	return hash(hash(hash(offset64, key1), key2), key3)
}

func hash4[A, B, C, D comparable](key1 A, key2 B, key3 C, key4 D) uint64 {
	return hash(hash(hash(hash(offset64, key1), key2), key3), key4)
}

func hash5[A, B, C, D, E comparable](key1 A, key2 B, key3 C, key4 D, key5 E) uint64 {
	return hash(hash(hash(hash(hash(offset64, key1), key2), key3), key4), key5)
}

func hash6[A, B, C, D, E, F comparable](key1 A, key2 B, key3 C, key4 D, key5 E, key6 F) uint64 {
	return hash(hash(hash(hash(hash(hash(offset64, key1), key2), key3), key4), key5), key6)
}

func hash7[A, B, C, D, E, F, G comparable](key1 A, key2 B, key3 C, key4 D, key5 E, key6 F, key7 G) uint64 {
	return hash(hash(hash(hash(hash(hash(hash(offset64, key1), key2), key3), key4), key5), key6), key7)
}

func hash[A comparable](hash uint64, key A) uint64 {
	switch v := any(key).(type) {
	case string:
		return hashString(hash, v)
	case int:
		return hashInt(hash, uint64(v))
	case int8:
		return hashInt(hash, uint64(v))
	case int16:
		return hashInt(hash, uint64(v))
	case int32:
		return hashInt(hash, uint64(v))
	case int64:
		return hashInt(hash, uint64(v))
	case uint:
		return hashUint(hash, uint64(v))
	case uint8:
		return hashUint(hash, uint64(v))
	case uint16:
		return hashUint(hash, uint64(v))
	case uint32:
		return hashUint(hash, uint64(v))
	case uint64:
		return hashUint(hash, v)
	case uintptr:
		return hashUint(hash, uint64(v))
	case float32:
		return hashFloat(hash, math.Float64bits(float64(v)))
	case float64:
		return hashFloat(hash, math.Float64bits(v))
	case bool:
		return hashBool(v)
	default:
		panic(fmt.Sprintf("unsupported type for caching %T", key))
	}
}

//func hashString(hash uint64, key string) uint64 {
//	for i := 0; i < len(key); i++ {
//		hash = (hash ^ uint64(key[i])) * prime64
//	}
//	return hash
//}

func hashString(hash uint64, key string) uint64 {
	length := len(key)

	// for loop unrolling
	// Process four characters at a time
	for i := 0; i < length/4*4; i += 4 {
		hash = (hash ^ uint64(key[i])) * prime64
		hash = (hash ^ uint64(key[i+1])) * prime64
		hash = (hash ^ uint64(key[i+2])) * prime64
		hash = (hash ^ uint64(key[i+3])) * prime64
	}
	// Process remaining characters
	for i := length / 4 * 4; i < length; i++ {
		hash = (hash ^ uint64(key[i])) * prime64
	}
	return hash
}
func hashInt(hash uint64, key uint64) uint64 {
	return (hash ^ key) * prime64
}

func hashUint(hash uint64, key uint64) uint64 {
	return (hash ^ key) * prime64
}

func hashFloat(hash uint64, key uint64) uint64 {
	return (hash ^ key) * prime64
}

func hashBool(key bool) uint64 {
	if key {
		return trueHash
	}
	return falseHash

}
