package main

import "sort"

func main() {

}

func maxChunksToSorted(arr []int) int {
	res := 0
	n := len(arr)
	m := make(map[int]int)
	temp := make([][2]int, n)
	for i := 0; i < n; i++ {
		m[arr[i]]++
		temp[i] = [2]int{arr[i], m[arr[i]]}
	}
	target := make([][2]int, n)
	copy(target, temp)
	sort.Slice(target, func(i, j int) bool {
		if target[i][0] == target[j][0] {
			return target[i][1] < target[j][1]
		}
		return target[i][0] < target[j][0]
	})
	cur := temp[0]
	for i := 0; i < n; i++ {
		if compare(cur, temp[i]) == true { // 小于temp[i]更新
			cur = temp[i]
		}
		if cur == target[i] {
			res++
		}
	}
	return res
}

func compare(a, b [2]int) bool {
	if a[0] == b[0] {
		return a[1] < b[1]
	}
	return a[0] < b[0]
}
