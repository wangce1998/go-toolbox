package util

import (
	"math/rand"
	"strconv"
	"time"
)

// RandomStr 生成指定长度的字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Int64ToStrings int64切片转string切片
func Int64ToStrings(arr []int64) []string {
	var strings []string
	for _, v := range arr {
		strings = append(strings, strconv.FormatInt(v, 10))
	}
	return strings
}

// IntToStrings int切片转string切片
func IntToStrings(arr []int) []string {
	var strings []string
	for _, v := range arr {
		strings = append(strings, strconv.Itoa(v))
	}
	return strings
}
