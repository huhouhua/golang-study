package _map

import (
	"fmt"
	"testing"
)

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"a": 1, "b": 4, "c": 9}
	for k, v := range m1 {
		fmt.Printf("k:%s v:%d \n", k, v)
	}
}
