package interview

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	m := 10
	defer fmt.Printf("first defer:%d \n", m)
	m = 100
	defer fmt.Printf("second defer: %d \n", m)
}

func funcDefer() (sum int) {
	sumA := 100
	sumB := 100
	sum = sumA + sumB
	defer func() {
		fmt.Printf("func2 first %d \n", sum)
	}()
	defer fmt.Printf("func2 second %d \n", sum)
	return sum + 10
}

func TestFuncDefer(t *testing.T) {
	funcDefer()
}
