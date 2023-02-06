package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WithTimeout()
	//WithCancel()
	WithDeadline()
}

func WithTimeout() {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)

	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("手动信号 5秒到了 完成！")
			return
		default:
			fmt.Println("default!")
		}
	}

}
func WithCancel() {

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for i := 4; i > 0; i-- {
			time.Sleep(time.Second)
			fmt.Printf("还剩%d秒 \n", i)
		}
		defer cancel()
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("收到信号 被取消！")
			return
		default:
			fmt.Println("等待中!")
			time.Sleep(time.Second)
		}
	}

}
func WithDeadline() {

	//创建一个子节点的context,3秒后自动取消
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*3))

	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("时间到了 完成！")
				return
			default:
				fmt.Println("等待中!")
				time.Sleep(time.Second)
			}
		}
	}()
	fmt.Println("开始等待5秒钟 等待ctx被自动取消！,time=", time.Now().Unix())
	time.Sleep(time.Second * 5)

}
