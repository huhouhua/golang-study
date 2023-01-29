package thread

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	for i := 0; i <= 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	time.Sleep(time.Millisecond * 50)
}

// 错误的写法
func TestInitError(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 50)
}

func doSomething(i int) {
	fmt.Print(i)
}

// 协程切换例子
func TestSwitchover(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			doSomething(0)

		}
	}()
	for {
		doSomething(1)
	}
}
