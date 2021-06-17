package main

import (
	"fmt"
	"sync"
)
// 锁
func main() {
	var mut sync.Mutex
	var gr sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		gr.Add(1)
		go func (i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			gr.Done()
		}(i)
	}
	gr.Wait()
	//time.Sleep(time.Second * 1)
	fmt.Printf("counter = %d", counter)
}

