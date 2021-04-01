package main

import (
	"fmt"
	"time"
)

//----------------------有关任务Task任务角色功能----------------------------

// 任务
type Task struct {
	f func() error
}

// 创建任务
func NewTask(arg_f func() error) *Task {
	return &Task{f: arg_f}
}

// 调用任务
func (t *Task) Execute() { t.f() }

//----------------------有关协程池 Pool任务角色功能----------------------------
type Pool struct {
	EntryChannel chan *Task
	JobsChannel  chan *Task
	MaxWokerNum  int
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task, 100),
		JobsChannel:  make(chan *Task, 100),
		MaxWokerNum:  cap,
	}
}

func (p *Pool) Woker(wokerId int) {
	// 内部管道中不断读取任务，并执行
	for task := range p.JobsChannel {
		task.Execute()
		fmt.Println("wokerID:", wokerId)
	}
}

func (p *Pool) Run() {
	// 1、创建worker
	for i := 0; i < p.MaxWokerNum; i++ {
		go p.Woker(i)
	}

	// 2、不断从入口获取任务
	go func() {
		for task := range p.EntryChannel {
			p.JobsChannel <- task
		}
	}()
}

//-----------------------------主程序------------------------------

func main() {
	// 1、创建任务
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	// 2、创建协程池
	p := NewPool(4)

	// 3、将任务传递给任务池 ，不断放入任务
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()
	p.Run()

	time.Sleep(time.Second * 30)
}

// 为什么需要协程池？
// 限定goroutine的个数，有效利用服务器cpu，内存资源，避免过多资源浪费在cpu切换上，一般会根据CPU个数来定
