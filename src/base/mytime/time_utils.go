package mytime

import "time"

// 获取当前时间毫秒时间戳
func NowUnixMilli() int64 {
	return time.Now().UnixNano() / 1e6
}

// 时间戳转时间对象
func UnixMilliToDateTime(unixMilli int64) time.Time {
	return time.Unix(0, unixMilli*1000000)
}
