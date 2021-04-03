package jobrunner

import (
	"github.com/Taoey/Lets-Go/third_part_lib/jobrunner/job"
	"github.com/bamzi/jobrunner"
	"testing"
	"time"
)

func Test00(t *testing.T) {
	jobrunner.Start(10, 1)                           //默认10个线程池
	jobrunner.Schedule("@every 2s", job.PrintTime{}) //每两秒执行一次
	time.Sleep(time.Second * 20)                     //防止主协程提前结束，导致定时任务关闭
}
