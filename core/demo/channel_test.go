package demo

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():   // 收到消息后  就可以取消了  所以如果有消息就是true
		return true
	default:		// 如果没收到消息  则需要等待  所以返回false
		return false
	}
}
//func cancel_1(can chan struct{}) {
//	can <- struct{}{}
//}
//
//func cancel_2(can chan struct{}) {
//	close(can)
//}

func TestCancel(t *testing.T) {
	//can := make(chan struct{}, 0)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for true {
				if isCancelled(ctx) {
					break
				}
			}
			fmt.Println(i, "Done")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}