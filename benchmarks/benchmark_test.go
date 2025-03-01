package benchmarks

import (
	M "github.com/AhmedGoudaa/go_memoize"
	memoize "github.com/emad-elsaid/memoize"
	"github.com/emad-elsaid/memoize/cache/adapters/patrickmn"
	"github.com/patrickmn/go-cache"
	"testing"
	"time"
)

func DoSomThingZero() string                                 { return "a" }
func DoSomThing1(a string) string                            { return a }
func DoSomThing2(a, b string) string                         { return a + b }
func DoSomThing3(a string, b string, c string) string        { return a + b + c }
func DoSomThing4(a string, b string, c string, s int) string { return a + b + c }

func BenchmarkDo0Mem(b *testing.B) {
	DoSomThingZeroMemoized := M.Memoize(DoSomThingZero, 10*time.Minute)

	for i := 0; i < b.N; i++ {
		DoSomThingZeroMemoized()
	}
}

func BenchmarkDo1Mem(b *testing.B) {
	DoSomThing1Memoized := M.Memoize1(DoSomThing1, 10*time.Minute)

	params := []string{"1111", "2222", "3333", "4444"}
	for i := 0; i < b.N; i++ {
		DoSomThing1Memoized(params[i%4])
	}
}
func BenchmarkDo1With_memoize(b *testing.B) {

	// from github.com/emad-elsaid/memoize
	// only for comparison
	// this memoize package is not used in the main code
	// and it only supports one argument functions
	var DoSomThing1Memoized = memoize.NewWithCache(
		patrickmn.GoCache[string](cache.New(10*time.Minute, 10*time.Minute)),
		DoSomThing1,
	)
	params := []string{"1111", "2222", "3333", "4444"}
	for i := 0; i < b.N; i++ {
		DoSomThing1Memoized(params[i%4])
	}
}

func BenchmarkDo2Mem(b *testing.B) {
	DoSomThing2Memoized := M.Memoize2(DoSomThing2, 10*time.Minute)

	params := []struct {
		a string
		b string
	}{
		{"1-", "1111"},
		{"2-", "2222"},
		{"3-", "3333"},
		{"4-", "4444"},
	}
	for i := 0; i < b.N; i++ {
		DoSomThing2Memoized(params[i%4].a, params[i%4].b)
	}
}

func BenchmarkDo3Mem(b *testing.B) {
	DoSomThing3Memoized := M.Memoize3(DoSomThing3, 10*time.Minute)

	params := []struct {
		a, b, c string
	}{
		{"1111", "2222", "3333"},
		{"4444", "5555", "6666"},
		{"7777", "8888", "9999"},
		{"aaaa", "bbbb", "cccc"},
	}
	for i := 0; i < b.N; i++ {
		DoSomThing3Memoized(params[i%4].a, params[i%4].b, params[i%4].c)
	}
}

func BenchmarkDo4Mem(b *testing.B) {
	DoSomThing4Memoized := M.Memoize4(DoSomThing4, 10*time.Minute)
	params := []struct {
		a, b, c string
		s       int
	}{
		{"1111", "2222", "3333", 1},
		{"4444", "5555", "6666", 2},
		{"7777", "8888", "9999", 3},
		{"aaaa", "bbbb", "cccc", 4},
	}
	for i := 0; i < b.N; i++ {
		DoSomThing4Memoized(params[i%4].a, params[i%4].b, params[i%4].c, params[i%4].s)
	}
}

// reflect version
func BenchmarkDo1MemRef(b *testing.B) {
	DoSomThing1Memoized := M.MemoizeRef(DoSomThing1, 10*time.Minute).(func(string) string)

	params := []string{"1111", "2222", "3333", "4444"}
	for i := 0; i < b.N; i++ {
		DoSomThing1Memoized(params[i%4])
	}
}

func BenchmarkDo2MemRef(b *testing.B) {
	DoSomThing2Memoized := M.MemoizeRef(DoSomThing2, 10*time.Minute).(func(string, string) string)

	params := []struct {
		a string
		b string
	}{
		{"1-", "1111"},
		{"2-", "2222"},
		{"3-", "3333"},
		{"4-", "4444"},
	}
	for i := 0; i < b.N; i++ {
		DoSomThing2Memoized(params[i%4].a, params[i%4].b)
	}
}

func BenchmarkDo3MemRef(b *testing.B) {
	DoSomThing3Memoized := M.MemoizeRef(DoSomThing3, 10*time.Minute).(func(string, string, string) string)

	params := []struct {
		a, b, c string
	}{
		{"1111", "2222", "3333"},
		{"4444", "5555", "6666"},
		{"7777", "8888", "9999"},
		{"aaaa", "bbbb", "cccc"},
	}
	for i := 0; i < b.N; i++ {
		DoSomThing3Memoized(params[i%4].a, params[i%4].b, params[i%4].c)
	}
}

func BenchmarkDo4MemRef(b *testing.B) {
	DoSomThing4Memoized := M.MemoizeRef(DoSomThing4, 10*time.Minute).(func(string, string, string, int) string)
	params := []struct {
		a, b, c string
		s       int
	}{
		{"1111", "2222", "3333", 1},
		{"4444", "5555", "6666", 2},
		{"7777", "8888", "9999", 3},
		{"aaaa", "bbbb", "cccc", 4},
	}
	for i := 0; i < b.N; i++ {
		DoSomThing4Memoized(params[i%4].a, params[i%4].b, params[i%4].c, params[i%4].s)
	}
}
