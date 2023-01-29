package main

import (
	"testing"
)

func TestVarParm(t *testing.T) {
	t.Log(Sum(1, 3, 3))
	t.Log(Sum(3, 5, 5))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("defer ")
	}()

	t.Log("panic")
	panic("异常")
}
