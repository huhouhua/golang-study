package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringBytes(t *testing.T) {
	s := "张"
	t.Log(len(s))

	e := "a"
	t.Log(len(e))
}

func TestStringToRune(t *testing.T) {
	s := "中"

	c := []rune(s)
	t.Log(len(c))

	t.Logf("unicode: %x", c[0])
	t.Logf("UTF8:%x", s)
}

func TestJoin(t *testing.T) {
	s := "中"
	b := "国"
	var str = s + b
	t.Log(str)

	c := fmt.Sprintf("%s %s", s, b)
	t.Log(c)

	t.Log(strings.Replace("oink oink oink", "k", "ky", 2))

	arrs := []string{s, b}

	t.Log(strings.Join(arrs, ""))
}
