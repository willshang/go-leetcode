package main

import "sort"

func main() {

}

func maxChunksToSorted(arr []int) int {
	res := 0
	n := len(arr)
	target := make([]int, n)
	copy(target, arr)
	sort.Ints(target)
	m := make(map[int]int)
	count := 0
	for i := 0; i < n; i++ {
		m[arr[i]]++
		if m[arr[i]] == 0 {
			count--
		} else if m[arr[i]] == 1 {
			count++
		}
		m[target[i]]--
		if m[target[i]] == 0 {
			count--
		} else if m[target[i]] == -1 {
			count++
		}
		if count == 0 {
			res++
		}
	}
	return res
}
