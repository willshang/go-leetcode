package main

import "sort"

func main() {

}

// leetcode791_自定义字符串排序
func customSortString(S string, T string) string {
	m := make(map[uint8]int)
	for i := 0; i < len(S); i++ {
		m[S[i]] = i
	}
	arr := []byte(T)
	sort.Slice(arr, func(i, j int) bool {
		return m[arr[i]] < m[arr[j]]
	})
	return string(arr)
}
