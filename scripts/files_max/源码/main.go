package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)
func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// 获取当前时间毫秒时间戳
func NowUnixMilli() int64 {
	return time.Now().UnixNano() / 1e9
}

func getDirList(dirpath string) ([]string, error) {
	var dir_list []string
	dir_err := filepath.Walk(dirpath,
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				dir_list = append(dir_list, path)
				return nil
			}

			return nil
		})
	return dir_list, dir_err
}

// 创建目标文件夹

func createTargetDir() string{
	milli := NowUnixMilli()
	target_name := "target_max_"+time.Unix(milli, 0).Format("20060102_150405")
	os.Mkdir(target_name,0777)              //创建目录
	return target_name
}

// 处理单个文件夹
func findMaxFile(path ,target_dir string) {

	files, _ := ioutil.ReadDir(path)
	sep := string(os.PathSeparator)

	max := 0
	max_filepath :=""
	max_filename :=""
	for _, f := range files {
		if !f.IsDir(){
			size,_ := strconv.Atoi(fmt.Sprintf("%v",f.Size()))
			if size > max{
				max = size
				max_filepath =  path +sep+f.Name()
				max_filename =  f.Name()

			}
		}
	}
	if max!=0 && max_filepath!=""{
		err :=File(max_filepath,target_dir+sep+max_filename)
		if err!=nil{
			fmt.Println(err)
		}
	}
}


func main() {
	var filename string
	fmt.Println("欢迎来到获取最大文件程序，请输入文件夹名称:")
	fmt.Scanln(&filename)

	//filename = "data"

	tar_dir := createTargetDir()
	time.Sleep(time.Second*1)
	dirs,_:= getDirList(filename)

	for _ , dir := range dirs{
		findMaxFile(dir,tar_dir)
		log.Print(dir+"处理完毕")
	}
	time.Sleep(time.Second*5)

}
