package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/tietang/props/v3/kvs"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

var WordExePath string
var WorkList *[]FileInfo

const (
	WORKLIST_DIR = "worklist"
)

func main() {

	// 判断记录文件夹存在性
	isExistWorklistDir()

	properties, _ := kvs.ReadPropertyFile("config.txt")
	WordExePath = properties.GetDefault("path", `C:\Program Files\Microsoft Office\root\Office16\WINWORD.EXE`)

	var filename string
	fmt.Println("请输入待处理文件夹名称:")
	fmt.Scanln(&filename)
	//filename = "data"

	// 判断记录存在性
	if !isExist("result_" + filename) {
		path := CreateTargetDir(filename)
		Dir(filename, path)

		WorkList = CreateWorkList("result_" + filename)

		dirs, _ := GetDirList(path)
		if len(dirs) <= 0 {
			return
		}
		RenameDirFiles(dirs[0])
	} else { //接着之前的记录进行工作
		WorkList = ReadWorkList("result_" + filename)
		RenameDirFilesByJson("result_" + filename)
	}

	fmt.Println("恭喜您，当前目录处理完毕")
	fmt.Println("当前窗口将于5秒后自动关闭...")
	time.Sleep(time.Second * 3)
}

// File copies a single file from src to dst
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

// Dir copies a whole directory recursively
func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

// 获取当前时间毫秒时间戳
func NowUnixMilli() int64 {
	return time.Now().UnixNano() / 1e9
}

// 获取文件夹文件列表
func GetDirList(dirpath string) ([]string, error) {
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
func CreateTargetDir(sourceFilepath string) string {
	//milli := NowUnixMilli()
	//target_name := "target_" + time.Unix(milli, 0).Format("20060102_150405")
	target_name := "result_" + sourceFilepath
	os.Mkdir(target_name, 0777) //创建目录
	return target_name
}

// 处理单个文件夹
func RenameDirFiles(path string) {
	files, _ := ioutil.ReadDir(path)
	sep := string(os.PathSeparator)
	for _, f := range files {
		if !f.IsDir() {
			absPath, _ := filepath.Abs(path + sep + f.Name())
			fmt.Println(absPath)
			OpenWordFile(absPath)
			fmt.Println(f.Name(), "是否修改文件名称（是:y 否：任意按键）")

			var renameFlag string
			fmt.Scanln(&renameFlag)
			if renameFlag == "y" {
				os.Rename(path+sep+f.Name(), path+sep+"！"+f.Name())
			}
			UpdateWorkList(path, f.Name())
		}
	}
}

func RenameDirFilesByJson(sourceDir string) {
	for _, file := range *WorkList {
		if file.Done == false {
			OpenWordFile(file.Filepath)
			fmt.Println(file.Filename, "是否修改文件名称（是:y 否：任意按键）")

			sep := string(os.PathSeparator)
			var renameFlag string
			fmt.Scanln(&renameFlag)
			if renameFlag == "y" {
				os.Rename(sourceDir+sep+file.Filename, sourceDir+sep+"！"+file.Filename)
			}
			UpdateWorkList(sourceDir, file.Filename)
		}
	}
}

// 获取某个文件夹文件绝对路径
func GetDirFileAbsPathsAndFilename(dir string) ([]string, []string) {
	resultFilepath := []string{}
	resultFilename := []string{}

	files, _ := ioutil.ReadDir(dir)
	sep := string(os.PathSeparator)
	for _, f := range files {
		if !f.IsDir() {
			absPath, _ := filepath.Abs(dir + sep + f.Name())
			resultFilepath = append(resultFilepath, absPath)
			resultFilename = append(resultFilename, f.Name())
		}
	}
	return resultFilepath, resultFilename
}

func OpenWordFile(path string) {
	Command(WordExePath, path)
}

func Command(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		fmt.Println("启动进程错误", err)
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

// 超时控制进程
func CommandWithTimeout(timeout time.Duration, name string, arg ...string) ([]byte, error) {
	ctxt, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	//fmt.Println("执行命令：%v %v", name, arg)
	cmd := exec.CommandContext(ctxt, name, arg...)

	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	if err := cmd.Start(); err != nil {
		fmt.Println("启动进程错误", err)
		return buf.Bytes(), err
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("进程异常中断或超时退出", err)
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

type FileInfo struct {
	Filepath string `json:"filepath"`
	Filename string `json:"filename"`
	Done     bool   `json:"done"`
}

// 创建工作进程
func CreateWorkList(sourceDir string) *[]FileInfo {
	paths, names := GetDirFileAbsPathsAndFilename(sourceDir)
	fileInfos := make([]FileInfo, len(paths))

	for i, path := range paths {
		fileInfos[i] = FileInfo{path, names[i], false}
	}

	// 打开文件句柄写入json
	filePtr, err := os.Create(WORKLIST_DIR + string(os.PathSeparator) + sourceDir + ".json")
	if err != nil {
		fmt.Println("文件创建失败", err.Error())
		return nil
	}

	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(fileInfos)
	if err != nil {
		fmt.Println("创建工作记录失败", err.Error())
	} else {
		fmt.Println("创建工作记录成功")
	}
	return &fileInfos
}

// 更新工作进程
func UpdateWorkList(sourceDir string, filename string) {
	curWorklist := *WorkList
	// 查找当前worklist
	for i, file := range curWorklist {
		if file.Filename == filename {
			tmpFile := FileInfo{
				file.Filepath,
				filename,
				true,
			}
			curWorklist[i] = tmpFile
			WorkList = &curWorklist
			break
		}
	}
	// 打开文件句柄写入json
	filePtr, err := os.Create(WORKLIST_DIR + string(os.PathSeparator) + sourceDir + ".json")
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	err = encoder.Encode(curWorklist)
	if err != nil {
		fmt.Println("更新进度失败", err.Error(), filename)
	} else {
		fmt.Println("更新进度成功", filename)
	}
}

// 从文件中读取json数据
func ReadWorkList(sourceDir string) *[]FileInfo {
	path := WORKLIST_DIR + string(os.PathSeparator) + sourceDir + ".json"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil
	}

	//读取的数据为json格式，需要进行解码
	files := []FileInfo{}
	err = json.Unmarshal(data, &files)
	if err != nil {
		return nil
	}
	return &files
}

func isExistWorklistDir() {
	temp_dir := WORKLIST_DIR
	_, err := os.Stat(temp_dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(temp_dir, os.ModePerm)
			if err != nil {
				fmt.Printf("mkdir failed![%v]\n", err)
			}
			return
		}
		fmt.Println("stat file error")
		return
	}
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
