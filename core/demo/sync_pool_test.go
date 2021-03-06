package demo

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	runtime.GC()
	//pool.Put(3)
	fmt.Println(v)
	v1, _ := pool.Get().(int)

	fmt.Println(v1)
	fmt.Println(v1)
	fmt.Println(v1)
}

func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	pool.Put(10)
	pool.Put(10)
	pool.Put(10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}