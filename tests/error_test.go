package main

import (
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log(err)
		}
	}()
	fmt.Println("dadads")
	panic("错误。。。。")
}
