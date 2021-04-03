package redigo

import (
	"fmt"
	"log"
	"testing"
)

// test :RedisKeyGet RedisKeySet
func Test01(t *testing.T) {
	bytes, e := RedisKeyGet("tao")
	log.Println(string(bytes), e)

	RediskeySet("tao", []byte("18"))

	bytes2, e := RedisKeyGet("tao")
	log.Println(string(bytes2), e)
}

// 使用redis-cell模块进行限流控制
// 通过reply数组中的值来控制函数的执行，可实现接口频率调用，以实现最终的限流效果
func TestRedisLimit(t *testing.T) {
	conn := Pool.Get()
	defer conn.Close()

	for i := 0; i < 20; i++ {
		reply, err := conn.Do("cl.throttle", "tao:reply1", 15, 30, 60, 1)
		fmt.Println(reply, err)
	}

}
