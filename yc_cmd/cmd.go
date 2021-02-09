package yc_cmd

import (
	"fmt"
	"time"
)

// windows窗口倒计时关闭
func Countdown(len int) {
	for i := len; i > 0; i-- {
		fmt.Printf("%d 后将关闭窗口……\n", i)
		time.Sleep(time.Second)
	}
}
