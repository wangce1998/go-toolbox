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

// StringIntersect 求两个字符串切片的交集
func StringIntersect(a []string, b []string) []string {
	var inter []string
	mp := make(map[string]bool)
	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}

// StringDiff 求两个字符串切片的差集
func StringDiff(a []string, b []string) []string {
	var diffArray []string
	temp := map[string]bool{}
	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = true
		}
	}
	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}
