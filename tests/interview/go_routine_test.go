package interview

import (
	"fmt"
	"testing"
	"time"
)

func TestGoRoutine(t *testing.T) {
	data := make(map[int]int, 10)
	for i := 1; i < 10; i++ {
		data[i] = i
	}
	for key, value := range data {
		key := key
		value := value
		go func() {
			fmt.Println("key ->", key, "v ->", value)
		}()
	}
	time.Sleep(time.Second * 5)
}
