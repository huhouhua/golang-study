package lock

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	var mutex = sync.Mutex{}
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			mutex.Lock()
			counter++
			defer mutex.Unlock()
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

func TestLockTest(t *testing.T) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*10)
	c := &contextLock{
		users: []*user{&user{}},
	}
	for i := 0; i < 100; i++ {
		go func(i int) {
			c.AddUser(i)
		}(i)
		for i := 0; i < 10; i++ {
			go func(i int) {
				c.AddUser(i)
			}(i)
		}
	}
	select {
	case <-ctx.Done():
		for _, u := range c.users {
			t.Logf("%v \n", u)
		}
	}
}

type contextLock struct {
	sync.Mutex
	users []*user
}

func (c *contextLock) AddUser(id int) {
	defer c.Unlock()
	c.Lock()

	for _, u := range c.users {
		if u.Id == id {
			return
		}
	}
	c.users = append(c.users, &user{
		Name: fmt.Sprintf("张三 %b", id),
		Id:   id,
	})
}

type user struct {
	Id   int
	Name string
}
