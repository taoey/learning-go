package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

type Food struct{}

func Test00(t *testing.T) {
	var num uint8
	num = 8
	i := reflect.TypeOf(num)
	fmt.Println(i) // uint8
}

func Test01(t *testing.T) {
	f := Food{}
	i := reflect.TypeOf(f)
	fmt.Println(i) // main_test.Food
}

func Test02(t *testing.T) {
	f := Food{}
	i := reflect.TypeOf(&f)
	fmt.Println(i) // *main_test.Food
}
