package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled2(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}
func TestCancel2(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx2 context.Context) {
			for {
				if isCancelled2(ctx2) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "任务被取消！")
		}(i, ctx)
	}
	cancelFunc()
	time.Sleep(time.Second * 1)
}
