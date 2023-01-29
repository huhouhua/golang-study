package main

import (
	"testing"
)

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(cap(s0), len(s0))
}

func TestSliceMake(t *testing.T) {
	arr := make([]int, 3, 5)
	t.Log(len(arr), cap(arr))
	t.Log(arr[0], arr[1], arr[2])

	arr = append(arr, 2)
	t.Log(len(arr), cap(arr))
	t.Log(arr[0], arr[1], arr[2], arr[3])
}

func TestCapacity(t *testing.T) {
	s := []int{}
	for i := 0; i <= 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceCompare(t *testing.T) {
	//a := []int{1, 2, 3, 4}
	//b := []int{2, 3, 4, 5}
	//if a == b {
	//
	//}

}
