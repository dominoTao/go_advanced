package demo


// 仅需 任意协程完成任务即可返回
// 完成所有任务  通过scp模式
import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)  // 通过buffer channel 来防止协程泄露
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <- ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second*1)
	t.Log("after:", runtime.NumGoroutine())

}


func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)  // 通过buffer channel 来防止协程泄露
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}