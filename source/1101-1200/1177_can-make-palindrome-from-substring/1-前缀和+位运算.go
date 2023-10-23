package main

import "math/bits"

func main() {

}

// leetcode1177_构建回文串检测
func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	arr := make([]int, n+1)
	cur := 0
	for i := 0; i < n; i++ {
		value := int(s[i] - 'a')
		cur = cur ^ (1 << value)
		arr[i+1] = cur
	}
	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		a, b, c := queries[i][0], queries[i][1], queries[i][2]
		status := arr[b+1] ^ arr[a]
		if bits.OnesCount(uint(status)) <= 2*c+1 { // 奇数次字母的个数，重新排列后最多允许2*c+1个
			res[i] = true
		}
	}
	return res
}
