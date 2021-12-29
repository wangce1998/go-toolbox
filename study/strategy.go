package study

import (
	"fmt"
)

// 策略模式 写入日志支持文件和数据库两种方式

type LogManager struct {
	Log
}

type Log interface {
	Write(string)
}

func NewLogManager(log Log) LogManager {
	return LogManager{log}
}

type File struct {
}

func (f File) Write(body string) {
	fmt.Println("写入文件日志:" + body)
}

type DB struct {
}

func (db DB) Write(body string) {
	fmt.Println("写入DB日志:" + body)
}
