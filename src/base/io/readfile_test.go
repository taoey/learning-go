package io

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

type student struct {
	Name string
	Age  int
}

//二维数组到student转化

// 传入一个指针，返回一个指针
func convert2Students(records *[][]string) *[]student {

	if len(*records) == 0 {
		return nil
	}

	result := make([]student, 0)
	// 转化records
	for index, row := range *records {
		if index < 1 {
			continue
		}
		name := row[0]
		age, _ := strconv.Atoi(row[1])
		result = append(result, student{name, age})
	}
	return &result
}

func Test_read(t *testing.T) {
	file, _ := ioutil.ReadFile("data1.txt")
	fmt.Println(string(file))
}

//使用官方库读取csv文件
func Test_read_csv(t *testing.T) {
	csvFile, _ := os.Open("data2.csv") //打开一个文件，获取到 *File
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, _ := csvReader.ReadAll()
	fmt.Println(records)
}

//读取csv文件并将数据转化为struct对象
func Test_read_csv_struts(t *testing.T) {
	csvFile, _ := os.Open("data2.csv") //打开一个文件，获取到 *File
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, _ := csvReader.ReadAll()

	result := convert2Students(&records)
	fmt.Println(*result) //通过在变量前加`*`来获取指针对应的值
}

//解决中文乱码问题 使用第三方库：golang.org/x/text/encoding
func Test_read_gbktring(t *testing.T) {
	file, _ := ioutil.ReadFile("GBK_data.txt")
	result := make([]byte, 3*len(file)/2+1) //设置目标byte数组，长度设置为原来的1.5倍

	decoder := simplifiedchinese.GBK.NewDecoder()             //设置转化器
	nDst, nSrc, err := decoder.Transform(result, file, false) //进行转化

	//结果查看
	fmt.Println(nDst, nSrc, err)
	fmt.Println(string(result), string(file))
	fmt.Println(len(file), len(result))

}
