package job

import (
	"fmt"
	"time"
)

// 打印时间任务 实现了Job接口
type PrintTime struct{}

func (e PrintTime) Run() {
	fmt.Println("time:", time.Now())
}
