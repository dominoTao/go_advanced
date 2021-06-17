package demo

import (
	"errors"
	"fmt"
	"testing"
	"time"
)
// 对象池
type ResuableObj struct {

}

type ObjPool struct {
	bufChan chan *ResuableObj    // 用于缓冲可重用对象
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//if err:=pool.ReleaseObj(&ResuableObj{}); err!=nil{
	//	t.Error(err)
	//}
	for i := 0; i < 11; i++ {
		if v,err:=pool.GetObj(time.Second*1);err!=nil {
			t.Error(err)
		}else{

			fmt.Printf("%T\n", v)
			if err:=pool.ReleaseObj(v); err!=nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}    //TODO  ???
	objPool.bufChan = make(chan *ResuableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan<-&ResuableObj{}
	}
	return &objPool				//TODO  ???
}

func (p *ObjPool)GetObj(timeout time.Duration) (*ResuableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool)ReleaseObj(obj *ResuableObj) error {
	select {
	case p.bufChan<-obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
