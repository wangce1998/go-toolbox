package xjson

import (
	"encoding/json"
)

// MarshalToString 转json字符串
func MarshalToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(b)
}

// ToStrings json字符串转字符串切片
func ToStrings(str string) []string {
	var s []string
	_ = json.Unmarshal([]byte(str), &s)
	
	return s
}
