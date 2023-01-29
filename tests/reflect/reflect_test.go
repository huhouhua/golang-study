package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.String:
		fmt.Println("是string类型！")
	case reflect.Float32, reflect.Float64:
		fmt.Println("是Float类型")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Int类型")
	default:
		fmt.Println("没有该类型！", t)
	}
}
func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}
func TestInvoke(t *testing.T) {
	var stu = &Student{
		Name: "张三",
		Age:  24,
	}
	t.Logf("Before: %#v", stu)
	Invoke(stu)
	t.Logf("After: %#v", stu)
}

func Invoke(v interface{}) {
	t := reflect.TypeOf(v)
	tagType := reflect.TypeOf(&Student{})
	if t.Kind() == tagType.Kind() {
		v := reflect.ValueOf(v)
		for i := 0; i < v.NumMethod(); i++ {
			method := t.Method(i)
			fmt.Printf("方法名称:%s \n", method.Name)
			var callPar []reflect.Value

			if method.Name == "SetAge" {
				callPar = append(callPar, reflect.ValueOf(2))
				method.Func.Call(callPar)
			} else {
				method.Func.Call(callPar)
			}
		}
		for i := 0; i < v.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("字段名称:%s  \n", field.Name)
		}
	}
}

type Student struct {
	Age  int
	Name string
}

func (s *Student) GetName() string {
	return s.Name
}
func (s *Student) SetAge(age int) {
	s.Age = age
}
