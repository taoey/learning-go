package redigo

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func init() {
	initLog()
	initRedis()
}


func initLog()  {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var (
	Pool *redis.Pool
)
// 初始化redis
func initRedis() {
	server := "localhost:6379"
	password :="tao"
	Pool = newPool(server,password)
	cleanupHook()
}

// 创建Redis连接池
func newPool(server string, pwd string) *redis.Pool {
	redisPool := redis.Pool{
		MaxIdle:     500,
		MaxActive:   2000,
		IdleTimeout: 60 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _, err := c.Do("AUTH", pwd); err != nil {
					errInfo := strings.ToLower(err.Error())
					log.Print(errInfo)
				}
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return &redisPool
}


func cleanupHook() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}