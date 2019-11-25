package goroutine

import (
    "sync"
    "runtime"
    "errors"
)

type Converter interface {
    run()
}

type GoroutinePool struct {
    closed bool
    sync.WaitGroup
    work chan Converter
}

func (p *GoroutinePool) Run(w Converter) error {
    if p.closed {
        return errors.New("channel is closed")
    }
    select {
    case p.work <- w:
        return nil
    default:
        return errors.New("channel is full")
    }
}

func (p *GoroutinePool) Close() {
    p.closed = true
    close(p.work)
    p.Wait()
}

func NewGoroutinePool() *GoroutinePool {
    p := GoroutinePool{
        work: make(chan Converter, 100),
    }
    size := runtime.NumCPU()-1
    p.Add(size)
    for i := 0; i < size; i++ {
        go func() {
            for w := range p.work {
                w.run()
            }
            p.Done()
        }()
    }
    return &p
}

type ConvertTask struct {
    CallBack       func()
}
func (t *ConvertTask) run() {
    t.CallBack()
}
func (t *ConvertTask) runTask(fun func()) {
    t.CallBack = fun
}