package pool

import (
	"errors"
	"fmt"
	"testing"
	"time"
	"unsafe"
)

type ReusableObj struct {
}

type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用的对象
}

// 创建对象
func NewObjPool(numOfObj int) *ObjPool {
	pool := ObjPool{}
	pool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		pool.bufChan <- &ReusableObj{}
	}
	return &pool
}

// 获取对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		fmt.Printf("%x 获取成功！\n ", unsafe.Pointer(ret))
		return ret, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}

// 归还对象
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		fmt.Printf("%x 归还成功！\n ", unsafe.Pointer(obj))
		return nil
	default: //归还不进去，返回异常
		return errors.New("overflow")
	}
}

func TestObjPoll(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if obj, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			if err := pool.ReleaseObj(obj); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}
