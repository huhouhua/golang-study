package cancel

import (
	"fmt"
	"sync"
	"testing"
)

// 生产数据
func dataProduct(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

// 消费数据
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			if data, isClose := <-ch; isClose {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	dataProduct(ch, &wg)
	dataReceiver(ch, &wg)
	wg.Wait()
}
