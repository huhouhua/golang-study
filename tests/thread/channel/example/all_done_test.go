package example

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask2(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is form %d", id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask2(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for i := 0; i < numOfRunner; i++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) //获取当前系统的协程数
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine()) //获取当前系统的协程数
}
