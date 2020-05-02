# Golang--超时控制



问题场景：我们有一个方法A，

- 当方法A正常执行时，大概2秒返回正常结果
- 方法A发生e1错误时，大概4秒返回错误信息
- 方法A发生e2错误时，大概60秒返回错误信息
- 方法A发生e3错误时，大概120秒返回错误信息



在进行对外接口的提供时，60秒，120秒这个时间太长了，我们可以把e2，e3错误列为`超时错误`。但是需要如何进行设计呢？



多说无益，来看代码



1、初级版程序

如下程序中，我们有一个名叫`slowOperation`的方法。我们在main中调用它，我们期望在5秒之内获得到结果，不能让程序一直处于阻塞状态，否则main中的其他程序无法执行。

```go
package main

import (
	"log"
	"time"
)
func init() {
	log.SetFlags(log.Ldate |log.LstdFlags| log.Lshortfile)
}

func main() {
	log.Println("Program started...")
	result := slowOperation()
	log.Println("Program finished...",result)

	// do other thing
}

func slowOperation() bool {
	log.Println("slowOperation started...")
	time.Sleep(60 * time.Second)
	log.Println("slowOperation finished...")
	return true
}
```



执行结果：

```
=== RUN   TestMainTest
2019/12/04 18:10:32 FuncWithTimeout_test.go:13: Program started...
2019/12/04 18:10:32 FuncWithTimeout_test.go:21: slowOperation started...
2019/12/04 18:11:32 FuncWithTimeout_test.go:23: slowOperation finished...
2019/12/04 18:11:32 FuncWithTimeout_test.go:15: Program finished... true
--- PASS: TestMainTest (60.01s)
PASS
```

我们可以看到整个程序执行了60s，并不是我们想要的结果，需要对程序进行优化处理。



2、改进版

改进思路：使用Golang自带的context包下的`WithTimeout()`函数，结合协程，通道进行控制

```go
package main

import (
	"context"
	"errors"
	"log"
	"time"
)

func init() {
	log.SetFlags(log.Ldate |log.LstdFlags| log.Lshortfile)
}
func  main() {
	log.Println("Program started...")
	err := operationWithTimeout(time.Second * 5)
	if err!=nil{
		log.Println(err)
	}
	log.Println("Program finished...")
}

// 新建方法：包装slowOperation，进行超时控制
func operationWithTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	c := make(chan bool)
	// 起一个协程执行slowOperation
	go func() {
		slowOperation(c)
	}()

	select {
	case <-ctx.Done():
		//log.Println(ctx.Err())
		return errors.New("send timeout")  // 返回自定义错误
	case result:= <-c:
		log.Println(result)
		return nil
	}
	log.Println("operationWithTimeout finished...")
	return nil
}

func slowOperation(c chan bool) {
	log.Println("slowOperation started...")
	time.Sleep(60 * time.Second)  // 模拟耗时程序， 其中可以包含 c <- false
	c <- true
	log.Println("slowOperation finished...")
}
```



执行结果：

```
=== RUN   TestMainTest
2019/12/04 18:15:03 FunWithTimeout2_test.go:15: Program started...
2019/12/04 18:15:03 FunWithTimeout2_test.go:47: slowOperation started...
2019/12/04 18:15:08 FunWithTimeout2_test.go:18: send timeout
2019/12/04 18:15:08 FunWithTimeout2_test.go:20: Program finished...
--- PASS: TestMainTest (5.01s)
PASS
```

我们可以看到整个过程只用了5秒钟



其他相关，大家可以去看一下golang的 net.http.client的源码，其中的setRequestCancel这个函数使用了类似的思路，另外为了加深理解，可以去看看context.context 的源码。