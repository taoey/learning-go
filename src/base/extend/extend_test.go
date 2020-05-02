package extend

import (
	"fmt"
	"testing"
)

// golang的多态

// ----------- 接口定义 -------------------------
type IAnimal interface {
	Eat() // 描述吃的行为
}

//------------ Animal ---------------------
// 动物 所有动物的父类
type Animal struct {
	Name string
}

// 动物的构造函数
func newAnimal(name string) *Animal {
	return &Animal{
		Name: name,
	}
}

// 动物去实现IAnimal中描述的吃的接口
func (a *Animal) Eat() {
	fmt.Printf("%v is eating food \n", a.Name)
}

//-------------- Cat -------------------
// 猫的结构体 组合了animal
type Cat struct {
	*Animal
}

// 实现猫的构造函数 初始化animal结构体
func newCat(name string) *Cat {
	return &Cat{
		Animal: newAnimal(name),
	}
}

//---------------- Dog ------------------
// 狗的结构体 组合了animal
type Dog struct {
	*Animal
}
// 狗的构造函数
func newDog(name string) *Dog {
	return &Dog{
		Animal:newAnimal(name),
	}
}

// 重写eat()方法
func (d *Dog) Eat()  {
	fmt.Printf("%v is eating bone \n", d.Name)
}

// --------------- Main --------------------------

func Test01(t *testing.T) {
	cat := newCat("eeee")
	cat.Eat()
}

func Test02(t *testing.T) {
	dog := newDog("hans")
	dog.Eat()
}

















