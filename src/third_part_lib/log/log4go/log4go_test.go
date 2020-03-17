package log4go

import (
	"github.com/jeanphorn/log4go"
	"os"
	"testing"
	"time"
)
var LOG log4go.Logger = nil

func setServerLogger() {
	LOG = make(log4go.Logger)
	lvl := log4go.INFO
	//switch CFG.UString("log.level", "INFO") {
	switch "INFO" {
	case "DEBUG":
		lvl = log4go.DEBUG
	case "INFO":
		lvl = log4go.INFO
	case "WARNING":
		lvl = log4go.WARNING
	case "ERROR":
		lvl = log4go.ERROR
	case "CRITICAL":
		lvl = log4go.CRITICAL
	default:
		lvl = log4go.INFO
	}
	LOG.AddFilter("stdout", lvl, log4go.NewConsoleLogWriter())
	if !FileOk("./log") {
		os.MkdirAll("./log", os.ModePerm)
	}
	logWriter := log4go.NewFileLogWriter("./log/run.log", true, true)
	logWriter.SetRotateSize(5 * 1024 * 1024)
	logWriter.SetRotateMaxBackup(7)
	LOG.AddFilter("file", lvl, logWriter)
}

func FileOk(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func Test01(t *testing.T) {
	setServerLogger()
	LOG.Info("hello","垃圾")
	time.Sleep(time.Second*100)
}