package ratelimit

import (
	"golang.org/x/time/rate"
	"testing"
)

// golang 准库限流器

func TestGolangLimit(t *testing.T) {

	// r : 代表每秒可以向 Token 桶中产生多少 token。Limit 实际上是 float64 的别名
	// b : Token 桶的容量大小。
	limiter := rate.NewLimiter(10, 1)

	limiter.Allow()

}
