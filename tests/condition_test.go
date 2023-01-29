package main

import (
	"strconv"
	"testing"
)

func TestWhileLoop(t *testing.T) {

	n := 0
	for n <= 10 {
		t.Log(n)
		n++
	}

}

func TestFor(t *testing.T) {
	n := 5
	for i := 0; i < n; i++ {
		t.Log(i)
	}
	users := []*User{}
	for i := 0; i < 5; i++ {
		users = append(users, &User{
			Name: "user" + strconv.Itoa(i),
		})
	}

	for _, user := range users {
		t.Logf("%v \n ", user.Name)
	}
}

func TestIf(t *testing.T) {
	if a := 2; a == 2 {
		t.Log(a)
	} else if a == 3 {
		t.Log(a)
	}

}

func TestSwitch(t *testing.T) {
	for i := 0; i <= 5; i++ {
		switch {
		case i == 0:
			t.Log(i)
		case i == 1:
			t.Log(i)
		case i == 2:
			t.Log(i)
		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i <= 5; i++ {
		switch i {
		case 0, 1:
			t.Log(i)
		case 2, 3:
			t.Log(i)
		default:
			t.Logf("defalt: %d \n", i)
		}

	}
}
