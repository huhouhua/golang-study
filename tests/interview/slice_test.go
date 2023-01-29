package interview

import (
	"fmt"
	"testing"
)

// 少于1024双倍扩容，大于1024是1.24倍扩容
//1024 * 1.25 = 1280
//1024 * 1.25 * 4 =5120   查表5376/4=1344
//1024 * 1.25 * 8 =10248  查表10248/8=1280

func TestSlice(t *testing.T) {
	var arr1 []int32
	for i := 0; i < 1025; i++ {
		arr1 = append(arr1, int32(i))
	}
	fmt.Printf("len=%d cap=%d", len(arr1), cap(arr1))
}
