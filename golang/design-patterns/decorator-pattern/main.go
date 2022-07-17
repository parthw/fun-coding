// Decorator pattern in golang
// Can be used in logging, instrumentation or caching.

package main

import (
	"fmt"
	"sync"
	"time"
)

type doFunc func(int) int

func square(i int) int {
	time.Sleep(2 * time.Second)
	return i * i
}

func wrap(do doFunc) doFunc {
	return func(i int) int {
		// Here we created another fn function because of defer
		fn := func(i int) (r int) {
			defer func(t time.Time) { fmt.Printf("Time taken = %v, result = %v", time.Since(t), r) }(time.Now())
			return do(i)
		}
		return fn(i)
	}
}

func wrapCache(do doFunc, cache *sync.Map) doFunc {
	return func(i int) int {
		v, ok := cache.Load(i)
		if ok {
			return v.(int)
		}
		r := do(i)
		cache.Store(i, r)
		return r
	}
}

func main() {
	fmt.Println(wrap(square)(10))

	cache := &sync.Map{}
	fmt.Println(wrap(wrapCache(square, cache))(10))
	fmt.Println(wrap(wrapCache(square, cache))(20))
	fmt.Println(wrap(wrapCache(square, cache))(10))
	fmt.Println(wrap(wrapCache(square, cache))(20))
}
