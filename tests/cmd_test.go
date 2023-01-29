package main

import (
	"fmt"
	"testing"
)

func TestFirstTry(t *testing.T) {
	var a int = 1
	var b int32 = 2
	c := int32(a) + b
	fmt.Printf("%T \n", c)
}

func TestString(t *testing.T) {

	var s string

	t.Logf("%s \n ", s)
}
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{2, 4, 5, 5}

	t.Log(a == b)
	for i := range b {
		t.Log(b[i])
	}
}

type User struct {
	Name string
}
