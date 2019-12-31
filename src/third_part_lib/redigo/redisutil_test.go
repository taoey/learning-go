package redigo

import (
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"testing"
)
// 可单独运行的redis连接demo
func TestConnect(t *testing.T) {
	//使用redis封装的Dial进行tcp连接
	dialOption := redis.DialPassword("tao")
	c, err := redis.Dial("tcp", "localhost:6379", dialOption)
	errCheck(err)

	defer c.Close()

	//对本次连接进行set操作
	_, setErr := c.Do("set", "test:hello", "hello redis")
	errCheck(setErr)

	//使用redis的string类型获取set的k/v信息
	r, getErr := redis.String(c.Do("get", "test:hello"))
	errCheck(getErr)
	log.Print(r)

}

func errCheck(err error) {
	if err != nil {
		log.Print("sorry,has some error:", err)
		os.Exit(-1)
	}
}

// 以下测试用例测试redigo_util中的函数
