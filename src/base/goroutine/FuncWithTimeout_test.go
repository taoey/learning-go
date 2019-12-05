package goroutine

import (
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.Ldate |log.LstdFlags| log.Lshortfile)
}

func TestMainTest(t *testing.T) {
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