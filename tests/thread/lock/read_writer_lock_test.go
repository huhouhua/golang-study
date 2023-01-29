package lock

import (
	"fmt"
	"testing"
)

func TestReadWriteExample(t *testing.T) {

	for i := 0; i < 5; i++ {
		go func() {
			defer func() {
				fmt.Println("释放1")
			}()
			defer func() {
				fmt.Println("释放2")
			}()
			defer func() {
				fmt.Println("释放3")
			}()
			defer func() {
				fmt.Println("释放4")
			}()
			defer func() {
				fmt.Println("释放5")
			}()
		}()
	}
}
