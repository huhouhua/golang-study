package string

import (
	"bytes"
	"strconv"
	"strings"
	"testing"
)

func TestConcatBytesBuffer(t *testing.T) {
	count := 10
	number := 5000
	for i := 0; i < count; i++ {
		var buf bytes.Buffer
		for j := 0; j < number; j++ {
			buf.WriteString(strconv.Itoa(j))
		}
		t.Log(buf.String())
	}
}
func TestConcatBuilder(t *testing.T) {
	count := 10
	number := 5000
	for i := 0; i < count; i++ {
		var buf strings.Builder
		for j := 0; j < number; j++ {
			buf.WriteString(strconv.Itoa(j))
		}
		t.Log(buf.String())
	}
}
