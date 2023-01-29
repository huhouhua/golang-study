package main

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
	"unsafe"
)

func TestInitObj(t *testing.T) {
	emp := new(Employee)
	emp.Id = "emp1"
	emp.Name = "lis"
	emp.Age = 44
	emp2 := Employee{Name: "e2", Id: "2", Age: 33}
	t.Logf("%v ", emp)
	t.Logf("%v ", emp2)
}

func TestPointer(t *testing.T) {
	e := &Employee{"0", "zhang", 22}
	t.Log(e.String())
	t.Log(e.String2())
	t.Log(unsafe.Pointer(&e.Name))
}

func (e Employee) String() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s  Name: %s Age:%d", e.Id, e.Name, e.Age)
}
func (e *Employee) String2() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s  Name: %s Age:%d", e.Id, e.Name, e.Age)
}

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCacheMap(t *testing.T) {
	cache := NewMapImpl()
	go func() {
		id := 0
		for id < 50 {
			if val, err := cache.Get("id"); err == nil {
				fmt.Println(val)
				id++
			}
		}
	}()
	for i := 0; i < 50; i++ {
		cache.Add("id", strconv.Itoa(i))
	}
	time.Sleep(time.Second * 3)
}

type IMap interface {
	Get(key string) (string, error)
	Add(key string, value string)
	Delete(key string) bool
}
type MapImpl struct {
	cacheMap sync.Map
}

func (m *MapImpl) Get(key string) (string, error) {
	//if _, ok := m.cacheMap[key]; ok {
	if val, ok := m.cacheMap.Load(key); ok {
		return val.(string), nil
	} else {
		return "", fmt.Errorf("id:%s not exist", key)
	}

}

func (m *MapImpl) Add(key string, value string) {
	m.cacheMap.Store(key, value)
}
func (m *MapImpl) Delete(key string) bool {
	if _, ok := m.cacheMap.Load(key); ok {
		m.cacheMap.Delete(key)
		return true
	}
	return false
}
func NewMapImpl() IMap {
	return &MapImpl{
		cacheMap: sync.Map{},
	}
}

func TestAnyStruct(t *testing.T) {
	//s1 := struct {
	//	age  int
	//	name string
	//}{
	//	name: "aa",
	//	age:  33,
	//}
	//s2 := struct {
	//	name string
	//	age  int
	//}{
	//	name: "aa",
	//	age:  33,
	//}
	//t.Log(s1 == s2)
}
