package goroutine

import (
	"fmt"
	"testing"
)

// Job类
type Job struct {
	id    int
	other []string
}

//20.320s
func Test00(t *testing.T) {
	jobList := make([]Job, 10000)
	for i := 0; i < 10000; i++ {
		temp := make([]string, 1024*12)
		jobList[i] = Job{i, temp}
	}

	for _, job := range jobList {
		fmt.Println(job)
	}
}

//21.307s
func Test01(t *testing.T) {

	jobList := make([]Job, 10000)
	for i := 0; i < 10000; i++ {
		temp := make([]string, 1024*12)
		jobList[i] = Job{i, temp}
	}

	jobs := make(chan Job)
	done := make(chan bool, len(jobList))

	// 推送协程
	go func(msg string) {
		//fmt.Println(msg,"执行了推送协程")
		for _, job := range jobList {
			//fmt.Println(msg,"加载了任务:",job)
			jobs <- job
		}
		close(jobs)
	}("puser")

	// 处理协程(加快此处处理速度)
	go func(msg string) {
		//fmt.Println(msg,"执行了处理协程")
		for job := range jobs {
			fmt.Println(msg, job) //模拟任务处理
			//time.Sleep(time.Second) //模拟执行时间
			done <- true // 处理完成，通知主协程
		}
	}("worker")

	// 阻塞主协程
	for i := 0; i < len(jobList); i++ {
		//fmt.Println("main","阻塞主协程")
		<-done // 阻塞等待接收
		//fmt.Println("main","任务",i,"执行完毕")
	}
}

//19.83s
func Test02(t *testing.T) {
	jobList := make([]Job, 10000)
	for i := 0; i < 10000; i++ {
		temp := make([]string, 1024*12)
		jobList[i] = Job{i, temp}
	}

	done := make(chan bool, len(jobList))

	go func() {
		for _, job := range jobList[:5000] {
			fmt.Println(job)
			done <- true
		}
	}()
	go func() {
		for _, job := range jobList[5000:] {
			fmt.Println(job)
			done <- true
		}
	}()

	// 阻塞主协程
	for i := 0; i < len(jobList); i++ {
		//fmt.Println("main","阻塞主协程")
		<-done // 阻塞等待接收
		//fmt.Println("main","任务",i,"执行完毕")
	}
}
