package channel

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "完成"
}

func asyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 2)
	go func() {
		ret := service()
		fmt.Println("返回结果.")
		retCh <- ret
		fmt.Println("服务退出！")
	}()
	return retCh
}

func TestSelectService(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100): //等待ret返回结果，超时为100毫秒
		t.Error("time out!")
	}
}
