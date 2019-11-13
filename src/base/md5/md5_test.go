package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

// 获取一个文件的MD5值
func Test01(t *testing.T) {
	dataStuffBytes, err := ioutil.ReadFile("0c1f58eff47f7e7ac21012cbad151da7_1502043.png")
	if err != nil {
		log.Print("error when open file", err)
	}
	stuffMd5Bytes := md5.Sum(dataStuffBytes)
	stuffMd5 := hex.EncodeToString(stuffMd5Bytes[:])

	fmt.Println(stuffMd5)
}