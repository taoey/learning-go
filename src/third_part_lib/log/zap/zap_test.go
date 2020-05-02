package zap

import (
	"go.uber.org/zap"
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
	//2019-09-26T14:22:31.914+0800	INFO	zap/zap_test.go:27	无法获取网址	{"url": "http://www.baidu.com", "attempt": 3, "backoff": "1s"}
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

	//2019-09-26T14:22:31.925+0800	DEBUG	zap/zap_test.go:33	hello
	logger.Debug("hello")
}
