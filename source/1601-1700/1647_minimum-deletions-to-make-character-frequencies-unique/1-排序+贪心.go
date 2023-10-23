package main

import "sort"

func main() {

}

// leetcode1647_字符频次唯一的最小删除次数
func minDeletions(s string) int {
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[int(s[i])]++
	}
	arr := make([]int, 0)
	for _, v := range m {
		arr = append(arr, v)
	}
	sort.Ints(arr)
	M := make(map[int]bool)
	res := 0
	for i := len(arr) - 1; i >= 0; i-- {
		if M[arr[i]] == false {
			M[arr[i]] = true
			continue
		}
		j := arr[i]
		for j >= 0 {
			if M[j] == false {
				M[j] = true
				res = res + arr[i] - j
				break
			}
			j--
		}
		if j == -1 {
			res = res + arr[i]
		}
	}
	return res
}
