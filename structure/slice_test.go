package structure

import (
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	u := &User{
		Name:     "huhouhua",
		Id:       12,
		LastName: "houhua",
	}
	_ = reflect.ValueOf(u).Elem()

}

type User struct {
	Name     string
	Id       int
	LastName string
}
