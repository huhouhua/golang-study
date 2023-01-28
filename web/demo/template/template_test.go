package template

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"html/template"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	tpl := template.New("Hello-World")
	tpl, err := tpl.Parse(`Hello {{ .Name}}`)

	require.NoError(t, err)

	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, User{Name: "huhouhua"})

	require.NoError(t, err)
	assert.Equal(t, `Hello huhouhua`, buffer.String())
}

func TestMapData(t *testing.T) {
	tpl := template.New("Hello-World")
	tpl, err := tpl.Parse(`Hello {{ .Name}}`)

	require.NoError(t, err)

	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, map[string]string{"Name": "huhouhua"})

	require.NoError(t, err)
	assert.Equal(t, `Hello huhouhua`, buffer.String())
}

func TestSlice(t *testing.T) {
	tpl := template.New("Hello-World")
	tpl, err := tpl.Parse(`Hello {{index . 0}}`)

	require.NoError(t, err)

	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, []string{"huhouhua"})

	require.NoError(t, err)
	assert.Equal(t, `Hello huhouhua`, buffer.String())
}

func TestBasic(t *testing.T) {
	tpl := template.New("Hello-World")
	tpl, err := tpl.Parse(`Hello {{.}}`)

	require.NoError(t, err)

	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, 123)

	require.NoError(t, err)
	assert.Equal(t, `Hello 123`, buffer.String())
}

func TestFuncCall(t *testing.T) {
	tpl := template.New("Hello-World")
	tpl, err := tpl.Parse(`Hello, {{.Hello "hu" "houhua"}}`)

	require.NoError(t, err)

	buffer := &bytes.Buffer{}
	err = tpl.Execute(buffer, FuncCall{})

	require.NoError(t, err)
	assert.Equal(t, `Hello, hu.houhua`, buffer.String())
}

type FuncCall struct {
}

func (f FuncCall) Hello(first string, last string) string {
	return fmt.Sprintf("%s.%s", first, last)
}

type User struct {
	Name string
}
