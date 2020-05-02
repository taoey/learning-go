package redigo

import (
	"log"
	"testing"
)

// test :RedisKeyGet RedisKeySet
func Test01(t *testing.T) {
	bytes, e := RedisKeyGet("tao")
	log.Println(string(bytes),e)

	RediskeySet("tao",[]byte("18"))

	bytes2, e := RedisKeyGet("tao")
	log.Println(string(bytes2),e)
}

