package _map

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 4, "c": 9}
	for k, v := range m1 {
		fmt.Printf("k:%s v:%d \n", k, v)
	}
}

func TestObjEqual(t *testing.T) {
	u1 := User{
		Name: "张三",
		Id:   123,
	}
	u2 := User{
		Name: "张三",
		Id:   123,
	}
	if reflect.ValueOf(u1) == reflect.ValueOf(u2) {
		t.Log("相等")
	} else {
		t.Log("不相等")
	}
}

type User struct {
	Id   int
	Name string
}
