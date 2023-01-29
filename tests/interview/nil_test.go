package interview

import (
	"fmt"
	"testing"
)

func TestNil(t *testing.T) {
	var x *int = nil
	var y interface{} = x
	fmt.Println(x == y)
	fmt.Println(x == nil)
	fmt.Println(y == nil)
}
