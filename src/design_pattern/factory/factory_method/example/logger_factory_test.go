package example

import (
	"testing"
)

func TestDbLoggerFactory(t *testing.T) {
	factory := DbLoggerFactory{}
	logger := factory.CreateLogger()
	logger.WriteLog()
}

func TestFileLoggerFactory(t *testing.T) {
	factory := FileLoggerFactory{}
	logger := factory.CreateLogger()
	logger.WriteLog()
}
