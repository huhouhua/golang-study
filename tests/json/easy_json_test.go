package json

import (
	"fmt"
	"testing"
)

var jsonStr2 = `{
	"basic_info":{
		"name":"Mike",
        "age":22
       },
    "job_info":{
       "skills":["C#","GO","Python"]
 	}
	}`

func TestEasyJson(t *testing.T) {

	e := Employee{}
	err := e.UnmarshalJSON([]byte(jsonStr2))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(e)
	if val, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(val))
	}
}

// 性能测试
func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr2))
		if err != nil {
			b.Error(err)
		}
		if _, err := e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
	b.StopTimer()
}
