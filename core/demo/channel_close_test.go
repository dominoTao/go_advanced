package demo

import (
	"fmt"
	"sync"
	"testing"
)

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)
	dataProducer(ch, &wg)
	dataReceiver(ch, &wg)
	wg.Wait()

}

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		//ch <- 1   //panic: send on closed channel
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <- ch ; ok{
				fmt.Println(data)
			}else {
				break
			}
		}
	}()
	wg.Done()
}