package main

import (
	"testing"
)

func TestArray(t *testing.T) {
	arr := [...]int{5, 3, 4, 5}

	for i := 0; i <= len(arr)-1; i++ {
		t.Log(arr[i])
	}
}
func TestMake(t *testing.T) {
	arr := make([]int, 2, 3)
	arr[0] = 51412
	arr[1] = 1124
	var ca = cap(arr)
	for i := 0; i <= ca; i++ {
		t.Log(arr[i])
	}
}

func TestArrayMulti(t *testing.T) {
	var arr_multi = [2][3]int{{1, 2, 3}, {2, 3, 5}}
	for i := 0; i <= cap(arr_multi)-1; i++ {
		for j := 0; j <= cap(arr_multi[i])-1; j++ {
			t.Log(arr_multi[i][j])
		}
	}
}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 3, 4, 4, 5}
	arr_section := arr[:]
	t.Log(arr_section)
}

// int: 256 * 2 * 8 =4096   4096/8 = 1184

// 33
func TestSlice(t *testing.T) {
	var arr []int64
	for i := 0; i < 513; i++ {
		arr = append(arr, int64(i))
	}
	t.Logf("len:%d cap:%d", len(arr), cap(arr))
}

func TestInitArray(t *testing.T) {
	arr := &[]int{1, 2, 3, 4}
	//initArray(arr)
	//t.Log(arr)
	initArraySecond(arr)
	t.Log(arr)

}

func initArray(arr []int) {
	arr = []int{4, 5, 6}

}
func initArraySecond(arr *[]int) {
	arr = &[]int{4, 5, 6}
}
