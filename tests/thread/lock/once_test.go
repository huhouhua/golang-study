package lock

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleton *Singleton
var once = sync.Once{}

func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Printf("create first %T.. \n", singleton)
		singleton = new(Singleton)
	})
	return singleton
}

func TestGetSingletonObj(t *testing.T) {
	var wg = sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x \n ", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
