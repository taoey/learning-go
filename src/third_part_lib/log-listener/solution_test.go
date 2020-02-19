package log_listener

import (
	"fmt"
	"github.com/taoey/go-log-listener/listener"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	logListener := listener.NewDefaultLogListener("access.log")
	logHandler := func(str string) interface{} {
		fmt.Print("日志解析：", str)
		return str
	}
	logStorageHandler := func(log interface{}) {
		fmt.Print("日志存储：", log)
	}
	logListener.SetHandler(logHandler, logStorageHandler)
	logListener.Run()

	time.Sleep(time.Minute * 3)
}
