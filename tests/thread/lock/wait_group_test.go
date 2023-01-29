package lock

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"testing"
)

func TestInitWaitGroup(t *testing.T) {
	var wg = sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1) //添加一个信号量
		go func() {
			//需要做的需事情
			wg.Done() //这个线程已经执行完
		}()
	}
	wg.Wait() // 等待所有的信号量全部执行完！

}

func TestWaitGroupExample(t *testing.T) {
	var mut = sync.Mutex{}
	var wg = sync.WaitGroup{}
	count := 0
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			count++
			wg.Done()

			fmt.Printf("协程ID:%d  被执行完。。。 \n", getId())
		}(i)
	}
	wg.Wait()
	fmt.Printf("所有的都线程被执行完:%d \n", count)
}

func getId() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
