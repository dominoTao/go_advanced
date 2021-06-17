package demo


//CSP 并发机制
import (
	"fmt"
	"testing"
	"time"
)

func TestAsyncService(t *testing.T) {
	AsyncService()
	otherTask()
	time.Sleep(time.Second * 5)
}

func TestSelect(t *testing.T) {
	select {
	case r := <- AsyncService():
		fmt.Println(r)
	case <-time.After(time.Millisecond*100):
		t.Error("timeout")
	}
}

func AsyncService() chan string{
	retCh := make(chan string, 1)
	go func() {
		fmt.Println("AsyncService 1.")
		retCh <- service()
		fmt.Printf("AsyncService 2. %v", <-retCh)
	}()
	return retCh
}

func otherTask() {
	fmt.Println("otherTask 1")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("otherTask 2")
}

func service() string {
	time.Sleep(time.Millisecond * 1)
	return "service this is the string for channel"
}