package redigo

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// 命令：get
func RedisKeyGet(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Print("error get key：", key, err)
		return data, fmt.Errorf("error getting key %s: %v", key, err)
	}
	log.Printf("success get key-value : %v-%v", key, string(data[:]))
	return data, err
}

// 命令：set
func RediskeySet(key string, value []byte) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		log.Print("error setting redis key: "+key, err)
	}
	log.Print("success setting redis key: " + key + " value: " + string(value[:]))
	return err
}

// 设置key的过期时间
func RedisExpire() {

}
