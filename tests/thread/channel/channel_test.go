package channel

import (
	"fmt"
	"testing"
	"time"
)

func Service() string {
	time.Sleep(time.Millisecond * 50)
	return "完成"
}

func OtherTask() {
	fmt.Println("任务进行中....")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("任务完成！")
}

func AsyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 2)
	go func() {
		ret := Service()
		fmt.Println("返回结果.")
		retCh <- ret
		fmt.Println("服务退出！")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	OtherTask()
	fmt.Println(<-retCh)
}

func TestService(t *testing.T) {
	fmt.Println(Service())
	OtherTask()
}
