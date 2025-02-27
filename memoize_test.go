package go_memoize

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMemoizeWithTTL(t *testing.T) {
	count := 0
	computeFn := func() int {
		count++
		return 1
	}
	memoizedFn := Memoize(computeFn, 1*time.Second)
	memoizedFn()
	memoizedFn()
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn()
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}

}

func TestMemoize1WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key int) int {
		count++
		return key * 2
	}
	memoizedFn := Memoize1(computeFn, 1*time.Second)
	memoizedFn(21)
	memoizedFn(21)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(21)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize2WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key1, key2 int) int {
		count++
		return key1 + key2
	}
	memoizedFn := Memoize2(computeFn, 1*time.Second)
	memoizedFn(20, 22)
	memoizedFn(20, 22)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(20, 22)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize3WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3 int) int {
		count++
		return key1 + key2 + key3
	}
	memoizedFn := Memoize3(computeFn, 1*time.Second)
	memoizedFn(10, 20, 12)
	memoizedFn(10, 20, 12)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(10, 20, 12)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize4WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4 int) int {
		count++
		return key1 + key2 + key3 + key4
	}
	memoizedFn := Memoize4(computeFn, 1*time.Second)
	memoizedFn(10, 10, 10, 12)
	memoizedFn(10, 10, 10, 12)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(10, 10, 10, 12)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}
func TestMemoizeWithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func() int {
		count++
		return 1
	}
	memoizedFn := Memoize(computeFn, 0)
	memoizedFn()
	memoizedFn()
	if count != 1 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize1WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key int) int {
		count++
		return key * 2
	}
	memoizedFn := Memoize1(computeFn, 0)
	memoizedFn(21)
	memoizedFn(21)
	if count != 1 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize2WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key1, key2 int) int {
		count++
		return key1 + key2
	}
	memoizedFn := Memoize2(computeFn, 0)
	memoizedFn(20, 22)
	memoizedFn(20, 22)
	if count != 1 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize3WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3 int) int {
		count++
		return key1 + key2 + key3
	}
	memoizedFn := Memoize3(computeFn, 0)
	memoizedFn(10, 20, 12)
	memoizedFn(10, 20, 12)
	memoizedFn(10, 20, 12)
	memoizedFn(10, 20, 12)
	if count != 1 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize4WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4 int) int {
		count++
		return key1 + key2 + key3 + key4
	}
	memoizedFn := Memoize4(computeFn, 0)
	memoizedFn(10, 10, 10, 12)
	memoizedFn(10, 10, 10, 12)
	memoizedFn(10, 11, 10, 12)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize5WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4, key5 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5
	}
	memoizedFn := Memoize5(computeFn, 1*time.Second)
	memoizedFn(1, 2, 3, 4, 5)
	memoizedFn(1, 2, 3, 4, 5)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(1, 2, 3, 4, 5)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize6WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4, key5, key6 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5 + key6
	}
	memoizedFn := Memoize6(computeFn, 1*time.Second)
	memoizedFn(1, 2, 3, 4, 5, 6)
	memoizedFn(1, 2, 3, 4, 5, 6)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(1, 2, 3, 4, 5, 6)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize7WithTTL(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4, key5, key6, key7 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5 + key6 + key7
	}
	memoizedFn := Memoize7(computeFn, 1*time.Second)
	memoizedFn(1, 2, 3, 4, 5, 6, 7)
	memoizedFn(1, 2, 3, 4, 5, 6, 7)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}

	time.Sleep(2 * time.Second)
	memoizedFn(1, 2, 3, 4, 5, 6, 7)
	if count != 2 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize5WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4, key5 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5
	}
	memoizedFn := Memoize5(computeFn, 0)
	memoizedFn(1, 2, 3, 4, 5)
	memoizedFn(1, 2, 3, 4, 5)
	if count != 1 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize6WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4, key5, key6 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5 + key6
	}
	memoizedFn := Memoize6(computeFn, 0)
	memoizedFn(1, 2, 3, 4, 5, 6)
	memoizedFn(1, 2, 3, 4, 5, 6)
	if count != 1 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoize7WithTTL_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(key1, key2, key3, key4, key5, key6, key7 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5 + key6 + key7
	}
	memoizedFn := Memoize7(computeFn, 0)
	memoizedFn(1, 2, 3, 4, 5, 6, 7)
	memoizedFn(1, 2, 3, 4, 5, 6, 7)
	memoizedFn(1, 2, 3, 4, 5, 44, 77)
	memoizedFn(1, 2, 3, 4, 5, 233, 1000)
	if count != 3 {
		t.Errorf("Expected 2, got %d", count)
	}
}

func TestMemoizeWithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func() int {
		atomic.AddInt32(&count, 1)
		return 1
	}
	memoizedFn := Memoize(computeFn, 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn()
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoize1WithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(key int) int {
		atomic.AddInt32(&count, 1)
		return key * 2
	}
	memoizedFn := Memoize1(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(21)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoize2WithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(key1, key2 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2
	}
	memoizedFn := Memoize2(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(20, 22)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoize3WithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(key1, key2, key3 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3
	}
	memoizedFn := Memoize3(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(10, 20, 12)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoize4WithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(key1, key2, key3, key4 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4
	}
	memoizedFn := Memoize4(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(10, 10, 10, 12)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoize5WithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(key1, key2, key3, key4, key5 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4 + key5
	}
	memoizedFn := Memoize5(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(1, 2, 3, 4, 5)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoize6WithTTL_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(key1, key2, key3, key4, key5, key6 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4 + key5 + key6
	}
	memoizedFn := Memoize6(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			memoizedFn(1, 2, 3, 4, 5, 6)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}
