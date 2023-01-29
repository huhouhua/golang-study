package cancel

import (
	"fmt"
	"testing"
	"time"
)

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T) {
	caneclChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "任务被取消！")
		}(i, caneclChan)
	}
	cancel_2(caneclChan)
	time.Sleep(time.Second * 1)
}
