package main

import (
	"github.com/robfig/cron/v3"
	"log"
)

// cron 表达式学习
func main() {
	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	select {}
}
