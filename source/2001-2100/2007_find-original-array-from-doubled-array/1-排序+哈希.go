package main

import "sort"

func main() {

}

// leetcode2007_从双倍数组中还原原数组
func findOriginalArray(changed []int) []int {
	res := make([]int, 0)
	n := len(changed)
	if n%2 == 1 {
		return nil
	}
	sort.Ints(changed)
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		value := changed[i]
		if m[value] == 0 { // 不是双倍的元素
			m[value*2]++ // 标记双倍
			res = append(res, value)
		} else {
			m[value]--
			if m[value] == 0 {
				delete(m, value)
			}
		}
	}
	if len(m) == 0 {
		return res
	}
	return nil
}
