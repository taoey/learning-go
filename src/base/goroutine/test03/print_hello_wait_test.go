package test03

import (
	"fmt"
	"sync"
	"testing"
)

// 不打印 "hello"，因为main goroutine 先结束执行，child goroutine来不及执行，就提前结束
func Test00(t *testing.T) {
	sayHello := func() {
		fmt.Println("hello")
	}
	go sayHello()
}

// 使用 sync.WaitGroup 保证 child goroutine 执行结束
func Test01(t *testing.T) {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

// 以下这个例子表明：goroutine 在它们所创建的相同地址空间内执行
func Test02(t *testing.T) {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation) // 执行结果：welcome
}

// 你猜这个函数打印什么？ 打印：三次 good day 想不到吧 [手动狗头]
func Test03(t *testing.T) {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

// 把test03修改为我们想要得到的效果
func Test04(t *testing.T) {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
