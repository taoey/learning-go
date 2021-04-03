package golang_queue

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"gopkg.in/eapache/queue.v1"
)

func TestSyncQueue_Pop(t *testing.T) {
	amount := 20
	queue := NewSyncQueue()
	//array :=   make([]int,amount)
	for i := 0; i < amount; i++ {
		aa := i
		fmt.Println(&aa, &i)
		go func() {
			//array[aa] = aa
			//fmt.Println(array)
			queue.Push(aa)
			fmt.Println("push", queue.buffer)
		}()

		//time.Sleep(time.Microsecond*19) // error    出现重复元素 {[0 2 2 3 4 5 6 7 8 9 10 11 12 13 14 15] 0 0 16}
		//time.Sleep(time.Microsecond*20) // pass
		//time.Sleep(time.Microsecond*25) // pass
		//time.Sleep(time.Microsecond*50) // pass
		//time.Sleep(time.Microsecond*100) // pass
		//time.Sleep(time.Microsecond*500) // pass
		//time.Sleep(time.Millisecond*1) // pass
		//time.Sleep(time.Millisecond*10) // pass
		//time.Sleep(time.Millisecond*100) //pass
		//time.Sleep(time.Second)  //pass

	}
	//for i:=0;i<amount/2;i++{
	//	go func() {
	//		queue.Remove()
	//		fmt.Println("remove",queue)
	//	}()
	//}
	time.Sleep(time.Minute)
}

func TestSyncQueue_TryPop(t *testing.T) {
	type fields struct {
		lock    sync.Mutex
		popable *sync.Cond
		buffer  *queue.Queue
		closed  bool
	}
	tests := []struct {
		name   string
		fields fields
		wantV  interface{}
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &SyncQueue{
				lock:    tt.fields.lock,
				popable: tt.fields.popable,
				buffer:  tt.fields.buffer,
				closed:  tt.fields.closed,
			}
			gotV, gotOk := q.TryPop()
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("SyncQueue.TryPop() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotOk != tt.wantOk {
				t.Errorf("SyncQueue.TryPop() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
