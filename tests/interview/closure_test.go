package interview

import (
	"fmt"
	"testing"
)

func func1() func(int) int {
	sum := 0
	return func(val int) int {
		sum += val
		return sum
	}
}
func printFunc1() {
	sumFunc := func1()
	fmt.Println(sumFunc(1))
	fmt.Println(sumFunc(1))
}

func func2() int {
	val := 10
	defer func() {
		val += 1
		fmt.Println(val)
	}()
	return val
}

func printFunc2() {
	fmt.Println(func2())
}

func TestFunc1(t *testing.T) {
	printFunc1()
}

func TestFunc2(t *testing.T) {
	printFunc2()
}
