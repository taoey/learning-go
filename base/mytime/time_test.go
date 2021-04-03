package mytime

import (
	"fmt"
	"testing"
	"time"
)

func Test01(t *testing.T) {
	time1 := "8:20"
	time2 := "9:20"
	t1, err1 := time.Parse("12:00", time1)
	t2, err2 := time.Parse("12:00", time2)

	fmt.Println(t1.Before(t2), err1, err2)
}

func Test02(t *testing.T) {
	time1 := "2015-03-20 8:50:29"
	time2 := "2015-03-21 09:04:25"
	//先把时间字符串格式化成相同的时间类型
	t1, err := time.Parse("2006-01-02 15:04:05", time1)
	t2, err := time.Parse("2006-01-02 15:04:05", time2)
	if err == nil && t1.Before(t2) {
		//处理逻辑
		fmt.Println("true")
	}
}

func Test03(t *testing.T) {
	y, m, d := time.Now().Date()
	fmt.Println(y, m, d)
}

func Test04(t *testing.T) {
	str := time.Now().Format("2006-01-02")
	fmt.Println(str)
}

func Test05(t *testing.T) {
	timer := time.After(time.Second * 5) // 和上面功能相同，使用After更为直观方便
	fmt.Println(<-timer)
}

func Test06(t *testing.T) {
	Timer := time.After(2 * time.Second)
	<-Timer //等价于NewTimer(d).C
	fmt.Println("timed out")
}

func Test07(t *testing.T) {
	now := NowUnixMilli()
	fmt.Println(now)
}

func Test08(t *testing.T) {
	hour, min, sec := time.Now().Clock()
	fmt.Println(hour, min, sec)
}

func Test09(t *testing.T) {
	dateTime := UnixMilliToDateTime(int64(1594191375000))
	fmt.Println(dateTime)
}
