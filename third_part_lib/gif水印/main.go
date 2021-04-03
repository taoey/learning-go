/**
 * @Author: taoey
 * @Description:
 * @File:  main
 * @Date: 2020/12/30 13:55
 */
package main

import (
	"flag"
	"github.com/issue9/watermark"
	"math/rand"
	"time"
)

func main() {
	var maskpath = flag.String("maskpath", "", "水印路径")
	var filepath = flag.String("filepath", "", "文件路径")
	flag.Parse()

	if *maskpath == "" || *filepath == "" {
		return
	}
	// 随机位置
	positions := []watermark.Pos{watermark.TopLeft, watermark.TopRight, watermark.BottomLeft, watermark.BottomRight}
	rand.Seed(time.Now().Unix())
	curPos := positions[rand.Intn(4)]

	// 在原图上添加水印
	w, err := watermark.New(*maskpath, 2, curPos)
	if err != nil {
		panic(err)
	}

	err = w.MarkFile(*filepath)
}
