package reflect

import (
	"errors"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "One", 2: "two", 3: "three"}
	b := map[int]string{1: "One", 2: "two", 3: "three"}
	t.Log("a == b?", reflect.DeepEqual(a, b))
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))
}

func TestFillNameAndAge(t *testing.T) {
	settingsMap := map[string]interface{}{"Name": "Mike", "Age": 40}
	emp := Employee{}
	if err := FillBySettings(&emp, settingsMap); err != nil {
		t.Fatal(err)
	}
	t.Log(emp)
	cus := new(Customer)
	if err := FillBySettings(cus, settingsMap); err != nil {
		t.Fatal(err)
	}
	t.Log(*cus)
}

func FillBySettings(str interface{}, settingsMap map[string]interface{}) error {
	//判断是否指针类型
	if reflect.TypeOf(str).Kind() != reflect.Ptr {
		//Elem() 获取指针指向的值
		if reflect.TypeOf(str).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to the struct type")
		}
	}
	if settingsMap == nil {
		return errors.New("settings is nil.")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range settingsMap {
		if field, ok = (reflect.ValueOf(str)).Elem().Type().FieldByName(k); !ok {

			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vObj := reflect.ValueOf(str)
			vObj = vObj.Elem()
			vObj.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return nil
}

type Employee struct {
	Name string
	Age  int
}

type Customer struct {
	Name string
	Age  int
}
