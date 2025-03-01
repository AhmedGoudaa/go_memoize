package go_memoize

import (
	"reflect"
	"time"
)

// MemoizeRef returns a memoized version of the compute function with a specified TTL.
func MemoizeRef(computeFn interface{}, ttl time.Duration) interface{} {
	cache := NewCache[uint64, interface{}](int64(ttl.Seconds()))
	fnValue := reflect.ValueOf(computeFn)
	fnType := fnValue.Type()

	return reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		// Create a hash of the arguments
		hash := hashArgs(args)
		// Get or compute the value
		result := cache.GetOrCompute(hash, func() interface{} {
			return fnValue.Call(args)[0].Interface()
		})
		return []reflect.Value{reflect.ValueOf(result)}
	}).Interface()
}

// hashArgs creates a hash from the function arguments using existing hashing functions.
func hashArgs(args []reflect.Value) uint64 {
	switch len(args) {
	case 1:
		return hash1(args[0].Interface())
	case 2:
		return hash2(args[0].Interface(), args[1].Interface())
	case 3:
		return hash3(args[0].Interface(), args[1].Interface(), args[2].Interface())
	case 4:
		return hash4(args[0].Interface(), args[1].Interface(), args[2].Interface(), args[3].Interface())
	case 5:
		return hash5(args[0].Interface(), args[1].Interface(), args[2].Interface(), args[3].Interface(), args[4].Interface())
	case 6:
		return hash6(args[0].Interface(), args[1].Interface(), args[2].Interface(), args[3].Interface(), args[4].Interface(), args[5].Interface())
	case 7:
		return hash7(args[0].Interface(), args[1].Interface(), args[2].Interface(), args[3].Interface(), args[4].Interface(), args[5].Interface(), args[6].Interface())
	default:
		panic("unsupported number of arguments")
	}
}
