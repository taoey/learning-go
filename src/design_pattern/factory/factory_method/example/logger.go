package example

import "fmt"

// 日志记录器接口
type Logger interface {
	WriteLog()
}

//----------------------数据库日志记录器-----------------------------
type DbLogger struct{}

func (this *DbLogger) WriteLog() {
	fmt.Println("write db log")
}

//----------------------文件日志记录器-----------------------------

type FileLogger struct{}

func (this *FileLogger) WriteLog() {
	fmt.Println("write file log")
}
