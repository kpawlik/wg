package wg

import (
	"sync"
	"sync/atomic"
)

type WaitGroup struct {
	wg  sync.WaitGroup
	mx  sync.Mutex
	max int32
	cnt atomic.Int32
}

func NewWaitGroup(max int) *WaitGroup {
	return &WaitGroup{
		wg:  sync.WaitGroup{},
		mx:  sync.Mutex{},
		max: int32(max),
		cnt: atomic.Int32{},
	}
}

func (w *WaitGroup) Add1() {
	if w.cnt.Load() >= w.max {
		w.wg.Wait()
	}
	w.wg.Add(1)
	w.cnt.Add(1)
}
func (w *WaitGroup) Done() {
	cnt := w.cnt.Load()
	cnt--
	w.cnt.Store(cnt)
	w.wg.Done()
}

func (w *WaitGroup) Wait() {
	w.wg.Wait()
}
