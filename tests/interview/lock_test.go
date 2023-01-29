package interview

import (
	"fmt"
	"sync"
	"testing"
)

var mu = sync.Mutex{}
var chain string

func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + "--> A"
	B()
}
func B() {
	chain = chain + "--> B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + "--> C"
}

func TestLock(t *testing.T) {
	chain = "lock"
	A()
	fmt.Println(chain)
}
