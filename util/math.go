package util

import (
	"math/rand"
)

// Subsets 求子集
// @see https://leetcode-cn.com/problems/power-set-lcci/
func Subsets(arr []interface{}) [][]interface{} {
	var res [][]interface{}
	res = append(res, []interface{}{})
	for _, v := range arr {
		for _, s := range res {
			sub := make([]interface{}, len(s))
			copy(sub, s)
			sub = append(sub, v)
			res = append(res, sub)
		}
	}
	return res
}

// RandInt 获取指定范围的随机数
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min+1) + min
}
