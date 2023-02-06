package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//WithTimeout()
	WithCancel()
}

func WithTimeout() {

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)

	defer cancel()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("5秒到了 完成！")
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
			fmt.Println("被取消！")
			return
		default:
			fmt.Println("等待中!")
			time.Sleep(time.Second)
		}
	}

}
