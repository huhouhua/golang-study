package lock

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	var mutex = sync.Mutex{}
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			mutex.Lock()
			counter++
			defer mutex.Unlock()
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}
