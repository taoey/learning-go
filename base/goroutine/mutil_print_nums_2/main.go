package main

import (
	"fmt"
)

func printNum(maxNum int, chanPre chan int, chanNext chan int, exit chan bool, id string) {
	for i := 0; i <= maxNum; i++ {
		num := <-chanPre
		select {
		case <-exit:
			return
		default:
			if num == maxNum+1 {
				exit <- true
			} else {
				fmt.Printf("id %s：%d\n", id, num)
				num += 1
				chanNext <- num
			}
		}
	}
}

func main() {

	exit := make(chan bool)
	maxNum := 20

	gNum := 5
	if gNum < 2 {
		return
	}
	chanArr := make([]chan int, gNum)

	for i := 0; i < gNum; i++ {
		chanArr[i] = make(chan int)
	}

	for i := 0; i < gNum; i++ {
		if i != len(chanArr)-1 {
			go printNum(maxNum, chanArr[i], chanArr[i+1], exit, string(rune(i)+'A'))
		} else {
			go printNum(maxNum, chanArr[i], chanArr[0], exit, string(rune(i)+'A'))
		}

	}
	chanArr[0] <- 0
	<-exit
}
