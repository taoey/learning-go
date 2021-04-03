package redis_sentinel

import (
	"fmt"
	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
	"log"
	"strings"
	"testing"
	"time"
)

var RedisConnPool *redis.Pool

func InitRedisSentinelConnPool() {
	redisAddr := "192.168.3.101:26479,192.168.3.101:26480,192.168.3.101:26481"
	redisAddrs := strings.Split(redisAddr, ",")
	masterName := "mymaster" // 根据redis集群具体配置设置

	//type dialOptions struct {
	//	readTimeout  time.Duration
	//	writeTimeout time.Duration
	//	dialer       *net.Dialer
	//	dial         func(network, addr string) (net.Conn, error)
	//	db           int
	//	password     string
	//	useTLS       bool
	//	skipVerify   bool
	//	tlsConfig    *tls.Config
	//}
	sntnl := &sentinel.Sentinel{
		Addrs:      redisAddrs,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {

			//c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			c, err := redis.Dial("tcp", addr, redis.DialPassword("smartrtb2win"))
			if err != nil {
				fmt.Println(err)
				return nil, err
			}
			return c, nil
		},
	}

	RedisConnPool = &redis.Pool{
		MaxIdle:     500,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntnl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: CheckRedisRole,
	}
}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("Role check failed")
	} else {
		return nil
	}
}

func Test01(t *testing.T) {
	InitRedisSentinelConnPool()

	conn := RedisConnPool.Get()
	defer conn.Close()
	for {
		reply, err := redis.String(conn.Do("get", "tao"))
		if err != nil {
			if err != redis.ErrNil {
				log.Println("error:", err)
			}

			fmt.Println("reply:", reply)
			break
		} // if
		time.Sleep(time.Second * 3)
	} // for
}
