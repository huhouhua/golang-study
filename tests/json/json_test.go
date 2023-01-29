package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
        "age":22
       },
    "job_info":{
       "skills":["C#","GO","Python"]
 	}
	}`

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}

// 性能测试
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {
		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err := json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}

// 反序列化对象
func TestUnmarshal(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	t.Log(*e)
}

// 序列化对象
func TestMarshal(t *testing.T) {
	e := &Employee{
		BasicInfo: BasicInfo{
			Name: "Mike",
			Age:  25,
		},
		JobInfo: JobInfo{
			Skills: []string{
				"GO", "C#", "Python",
			},
		},
	}
	if val, err := json.Marshal(e); err == nil {
		fmt.Println(string(val))
	} else {
		t.Error(err)
	}
}
