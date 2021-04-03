package factory_method

import (
	"fmt"
	"testing"
)

// 工厂方法模式

// 产品接口
type Product interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type Factory interface {
	Create() Product
}

// 具体产品类
type ConcreteProduct struct {
	a, b int
}

func (this *ConcreteProduct) SetA(a int) {
	this.a = a
}

func (this *ConcreteProduct) SetB(b int) {
	this.b = b
}

func (this *ConcreteProduct) Result() int {
	return this.a + this.b
}

// 具体工厂类
type ConcreteFactory struct{}

func (this *ConcreteFactory) Create() Product {
	return &ConcreteProduct{}
}

// 测试
func TestFactoryMethod(t *testing.T) {
	factory := ConcreteFactory{} // 通过配置文件获取

	product := factory.Create()
	product.SetA(1)
	product.SetB(2)

	result := product.Result()

	fmt.Println(result)
}
