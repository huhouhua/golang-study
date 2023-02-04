package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	WithTimeout()

}

func WithTimeout() {

	context, cancel := context.WithTimeout(context.TODO(), time.Second*5)

	defer cancel()
	for {
		select {
		case <-context.Done():
			fmt.Println("完成！")
			return
		default:
			fmt.Println("default!")
		}
	}

}
