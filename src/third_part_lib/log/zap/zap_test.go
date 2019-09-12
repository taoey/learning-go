package zap

import (
	"go.uber.org/zap"
	"log"
	"testing"
	"time"
)

func Test00(t *testing.T) {
	url := "www.mi.com"
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}

func Test01(t *testing.T) {
	// zap.NewDevelopment 格式化输出
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	logger.Info("无法获取网址",
		zap.String("url", "http://www.baidu.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

func Test02(t *testing.T) {
	log.Println("hello")
}
