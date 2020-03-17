package go_unarr

import (
	"fmt"
	"github.com/gen2brain/go-unarr"
	"testing"
)

func Test01(t *testing.T) {
	a, err := unarr.NewArchive("01.7z")
	if err != nil {
		panic(err)
	}
	defer a.Close()

	err2 := a.Extract("./")
	if err2 != nil {
		panic(err2)
	}
}

// 解压7z文件
// dest 目录如果不存在，则会进行新建
func Z7Extract(src,dest string) ([]string,error) {
	archive, err := unarr.NewArchive(src)
	defer archive.Close()

	if err !=nil{
		return nil,err
	}
	contents, err := archive.List()

	extractErr := archive.Extract(dest)
	if extractErr!=nil{
		return nil,extractErr
	}
	return contents,nil
}

func TestZ7Extract(t *testing.T) {
	src := "01.7z"
	dest := "./"
	contents, err := Z7Extract(src, dest)
	fmt.Println(contents,err)
}