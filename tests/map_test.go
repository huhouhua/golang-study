package main

import (
	"strconv"
	"testing"
)

func TestInitMap(t *testing.T) {
	map1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(map1[2])
	t.Logf("map1 length=%d \n", len(map1))

	map2 := map[int]int{}
	map2[4] = 16
	t.Logf("map2 length=%d \n", len(map2))

	map3 := make(map[int]int, 10)
	t.Logf("map3  length=%d \n ", len(map3))

}

func TestMapExist(t *testing.T) {
	m := map[int]int{1: 1, 2: 3}
	val := m[0]
	t.Log(val)

	////循环遍历
	//for key, val := range m {
	//	t.Logf(" key: %s  value: %s \n", key, val)
	//}
}

func TestMapFactory(t *testing.T) {
	factMap := map[int]func(key int) string{}
	for i := 0; i < 10; i++ {
		factMap[i] = func(key int) string {
			return strconv.Itoa(key * 2)
		}
	}

	for key, fact := range factMap {
		t.Log(fact(key))
	}
}

func TestMapForSet(t *testing.T) {
	mapSet := map[int]bool{}
	n := 1
	mapSet[n] = true
	if mapSet[n] {
		t.Logf("%d is exist ", n)
	} else {
		t.Logf("%d is not exist ", n)
	}
}
