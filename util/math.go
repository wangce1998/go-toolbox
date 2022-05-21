package util

// subsets 求子集
// @see https://leetcode-cn.com/problems/power-set-lcci/
func subsets(arr []interface{}) [][]interface{} {
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
