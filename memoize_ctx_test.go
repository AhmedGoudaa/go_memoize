package go_memoize

import (
	"context"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestMemoizeCtx_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context) int {
		count++
		return 1
	}
	memoizedFn := MemoizeCtx(computeFn, 0)
	memoizedFn(context.Background())
	memoizedFn(context.Background())
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx1_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key int) int {
		count++
		return key * 2
	}
	memoizedFn := MemoizeCtx1(computeFn, 0)
	memoizedFn(context.Background(), 21)
	memoizedFn(context.Background(), 21)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx2_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key1, key2 int) int {
		count++
		return key1 + key2
	}
	memoizedFn := MemoizeCtx2(computeFn, 0)
	memoizedFn(context.Background(), 20, 22)
	memoizedFn(context.Background(), 20, 22)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx3_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key1, key2, key3 int) int {
		count++
		return key1 + key2 + key3
	}
	memoizedFn := MemoizeCtx3(computeFn, 0)
	memoizedFn(context.Background(), 10, 20, 12)
	memoizedFn(context.Background(), 10, 20, 12)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx4_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key1, key2, key3, key4 int) int {
		count++
		return key1 + key2 + key3 + key4
	}
	memoizedFn := MemoizeCtx4(computeFn, 0)
	memoizedFn(context.Background(), 10, 10, 10, 12)
	memoizedFn(context.Background(), 10, 10, 10, 12)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx5_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key1, key2, key3, key4, key5 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5
	}
	memoizedFn := MemoizeCtx5(computeFn, 0)
	memoizedFn(context.Background(), 1, 2, 3, 4, 5)
	memoizedFn(context.Background(), 1, 2, 3, 4, 5)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx6_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key1, key2, key3, key4, key5, key6 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5 + key6
	}
	memoizedFn := MemoizeCtx6(computeFn, 0)
	memoizedFn(context.Background(), 1, 2, 3, 4, 5, 6)
	memoizedFn(context.Background(), 1, 2, 3, 4, 5, 6)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx7_NoExpiry(t *testing.T) {
	count := 0
	computeFn := func(ctx context.Context, key1, key2, key3, key4, key5, key6, key7 int) int {
		count++
		return key1 + key2 + key3 + key4 + key5 + key6 + key7
	}
	memoizedFn := MemoizeCtx7(computeFn, 0)
	memoizedFn(context.Background(), 1, 2, 3, 4, 5, 6, 7)
	memoizedFn(context.Background(), 1, 2, 3, 4, 5, 6, 7)
	if count != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context) int {
		atomic.AddInt32(&count, 1)
		return 1
	}
	memoizedFn := MemoizeCtx(computeFn, 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(context.Background())
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx1_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key int) int {
		atomic.AddInt32(&count, 1)
		return key * 2
	}
	memoizedFn := MemoizeCtx1(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 21)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx2_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key1, key2 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2
	}
	memoizedFn := MemoizeCtx2(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 20, 22)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx3_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key1, key2, key3 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3
	}
	memoizedFn := MemoizeCtx3(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 10, 20, 12)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx4_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key1, key2, key3, key4 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4
	}
	memoizedFn := MemoizeCtx4(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 10, 10, 10, 12)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx5_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key1, key2, key3, key4, key5 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4 + key5
	}
	memoizedFn := MemoizeCtx5(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 1, 2, 3, 4, 5)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx6_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key1, key2, key3, key4, key5, key6 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4 + key5 + key6
	}
	memoizedFn := MemoizeCtx6(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 1, 2, 3, 4, 5, 6)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}

func TestMemoizeCtx7_ConcurrentAccess(t *testing.T) {
	var count int32
	computeFn := func(ctx context.Context, key1, key2, key3, key4, key5, key6, key7 int) int {
		atomic.AddInt32(&count, 1)
		return key1 + key2 + key3 + key4 + key5 + key6 + key7
	}
	memoizedFn := MemoizeCtx7(computeFn, 1*time.Minute)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			memoizedFn(context.Background(), 1, 2, 3, 4, 5, 6, 7)
		}()
	}
	wg.Wait()
	if atomic.LoadInt32(&count) != 1 {
		t.Errorf("Expected 1, got %d", count)
	}
}
