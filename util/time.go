package util

import (
	"time"
)

// Datetime 获取完整时间
func Datetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Millisecond 获取当前毫秒
func Millisecond() int64 {
	return time.Now().UnixNano() / 1000000
}
