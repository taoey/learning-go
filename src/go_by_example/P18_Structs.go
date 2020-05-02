package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	tom := person{name: "tom", age: 12}

	fmt.Println(tom)
	fmt.Println(&tom)

	fmt.Println(tom.age)
	tom.age = 44
	fmt.Println(tom.age)
}
