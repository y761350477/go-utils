package time

import (
	"math/rand"
	"time"
)

// 获取日期时间
func NowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取日期
func NowDate() string {
	return time.Now().Format("2006-01-02")
}

// 获取当前时间
func NowTime() string {
	return time.Now().Format("15:04:05")
}

// 指定睡眠时间
func SleepSecond(i int) {
	time.Sleep(time.Duration(i) * time.Second)
}

// 指定随机区间睡眠时间
func RandSleepSecond(min, max int64) {
	if min >= max || min == 0 || max == 0 {
		time.Sleep(time.Duration(max) * time.Second)
	}

	time.Sleep(time.Duration(rand.Int63n(max-min)+min) * time.Second)
}
