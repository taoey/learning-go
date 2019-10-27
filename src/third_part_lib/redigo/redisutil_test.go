package redigo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"os"
)

func main() {
	//使用redis封装的Dial进行tcp连接
	dialOption := redis.DialPassword("tao")
	c, err := redis.Dial("tcp", "localhost:6379", dialOption)
	errCheck(err)

	defer c.Close()

	//对本次连接进行set操作
	_, setErr := c.Do("set", "startdemo", "hello redis")
	errCheck(setErr)

	//使用redis的string类型获取set的k/v信息
	r, getErr := redis.String(c.Do("get", "startdemo"))
	errCheck(getErr)
	fmt.Println(r)

}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
