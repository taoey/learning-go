package eventbus

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"testing"
	"time"
)

// https://github.com/asaskevich/EventBus

func calculator(a int, b int) {
	fmt.Printf("%d\n", a + b)
}

func slowCalculator(a, b int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("%d\n", a + b)
}

func Test01(t *testing.T) {
	bus := EventBus.New();
	bus.Subscribe("main:calculator", calculator); 		// 事件订阅
	bus.Publish("main:calculator", 20, 40); 		// 事件发布
	bus.Unsubscribe("main:calculator", calculator);	// 取消订阅
}

func Test02(t *testing.T) {
	bus := EventBus.New()
	bus.SubscribeAsync("main:slow_calculator", slowCalculator, false)

	bus.Publish("main:slow_calculator", 20, 60)

	fmt.Println("start: do some stuff while waiting for a result")
	fmt.Println("end: do some stuff while waiting for a result")

	bus.WaitAsync() // wait for all async callbacks to complete

	fmt.Println("do some stuff after waiting for result")
}