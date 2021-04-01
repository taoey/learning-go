package main

import (
	"fmt"
	"testing"
)

func TestNewManagerConcurrencyProblem(t *testing.T) {
	m1 := NewManagerConcurrencyProblem(1)
	m2 := NewManagerConcurrencyProblem(2)
	fmt.Println(m1 == m2)
}

func TestNewManagerConcurrency(t *testing.T) {
	m1 := NewManagerConcurrency(1)
	m2 := NewManagerConcurrency(2)
	fmt.Println(m1 == m2)
}
