package main

import (
	"fmt"
	"testing"
)

func TestProgrammer(t *testing.T) {
	var pro IProgrammer = &Programmer{}
	t.Log(pro.Write("张三"))
}

func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("222")
}

func TestEmptyInterfaceProgrammer(t *testing.T) {
	var pro = &Programmer{}
	DoProgrammer(pro)

}

func DoSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("int ", i)
	}
	if s, ok := p.(string); ok {
		fmt.Println("string", s)
	}
}

func DoProgrammer(p interface {
	IProgrammer
	ICode
}) {
	if i, ok := p.(IProgrammer); ok {
		fmt.Println("IProgrammer  ", i.Write("1"))
	}
	if i, ok := p.(ICode); ok {
		fmt.Println("IProgrammer  ", i.Read("1"))
	}
}

type ICode interface {
	Read(str string) string
}

type IProgrammer interface {
	Write(str string) string
}

type Programmer struct {
}

func (p Programmer) Read(str string) string {
	return fmt.Sprintf("read: %s ", str)
}

func (p Programmer) Write(str string) string {
	return fmt.Sprintf("write: %s ", str)
}
