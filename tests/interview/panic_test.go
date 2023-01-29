package interview

import (
	"fmt"
	"testing"
	"time"
)

func TestPanic(t *testing.T) {
	go func() {
		t := time.NewTicker(time.Second)
		for {
			select {
			case <-t.C:
				go func() {
					defer func() {
						if err := recover(); err != nil {
							fmt.Println(err)
						}
					}()
					fmt.Println(time.Now())
					proc()
				}()
			}
		}
	}()
	select {}
}

func proc() {
	panic("OK")
}
