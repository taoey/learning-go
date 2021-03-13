package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "os"
    "path"
    "path/filepath"
    "strings"
    "time"
)
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
    target_name := "target_"+time.Unix(milli, 0).Format("20060102_150405")
    os.Mkdir(target_name,0777)              //创建目录
    return target_name
}

// 处理单个文件夹
func renameDirFiles(path string) {

    if len(path)<=23{
        return
    }

    tarDate := strings.Replace(path[23:32],".","",-1)

    files, _ := ioutil.ReadDir(path)
    sep := string(os.PathSeparator)
    for _, f := range files {
        if !f.IsDir(){
            os.Rename(path +sep + f.Name(), path +sep + tarDate + f.Name())
        }
    }
}


func main() {
    var filename string
    fmt.Println("请输入文件夹名称:")
    fmt.Scanln(&filename)

    //filename = "files"

    path := createTargetDir()
    Dir(filename,path)
    dirs,_:= getDirList(path)

    for _ , dir := range dirs{
        renameDirFiles(dir)
        log.Print(dir+"处理完毕")
    }
    time.Sleep(time.Second*5)
}
