package file

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	f, err := os.Open("../../../assets/file/file.txt")
	require.NoError(t, err)

	data := make([]byte, 64)

	n, err := f.Read(data)
	fmt.Println(n)
	require.NoError(t, err)

	n, err = f.WriteString("hello")
	fmt.Println(n)

	fmt.Println(err)
	require.Error(t, err)
	f.Close()

}

func TestFile_Write(t *testing.T) {
	f, err := os.OpenFile("../../../assets/file/file.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	require.NoError(t, err)
	n, err := f.WriteString("hello")
	fmt.Println(n)
	f.Close()
}

func TestFile_Create(t *testing.T) {
	f, err := os.Create("../../../assets/file/file_copy.txt")
	require.NoError(t, err)
	n, err := f.WriteString("hello,world")
	fmt.Println(n)
	f.Close()
}

func TestWorkPath(t *testing.T) {
	dir, err := os.Getwd()
	require.NoError(t, err)
	fmt.Println(dir)
}
