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
