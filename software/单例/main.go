package main

import (
	"fmt"
	"sync"
)

var m *Manager

type Manager struct {
	Age int
}

func (m *Manager) GetAge() {
	fmt.Println(m.Age)
}

//--------------------------------

// 实例1：
// 并发情况下无法保证单例效果
func NewManagerConcurrencyProblem(age int) *Manager {
	// goroutine_2
	if m == nil {
		// goroutine_1
		m = &Manager{age}
		return m
	}
	return m
}

// 实例2：
// 并发可以效率低下
var lock sync.Mutex

func NewManagerConcurrency(age int) *Manager {
	lock.Lock()
	defer lock.Unlock()
	if m == nil {
		m = &Manager{age}
		return m
	}
	return m
}

var lockPref sync.Mutex

// 实例3
// 双重锁机制优化，提高效率
func NewManagerConcurrencyPref(age int) *Manager {
	if m == nil {
		lockPref.Lock()
		defer lockPref.Unlock()
		m = &Manager{age}
		return m
	}
	return m
}

// 实例4
// 利用go语言once只执行一次的特性，优雅的实现单例
var once sync.Once

func NewManagerOnce() *Manager {
	once.Do(func() {
		m = &Manager{}
	})
	return m
}
