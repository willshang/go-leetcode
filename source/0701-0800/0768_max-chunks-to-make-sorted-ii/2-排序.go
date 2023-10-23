package main

import "sort"

func main() {

}

// leetcode768_最多能完成排序的块II
func maxChunksToSorted(arr []int) int {
	res := 0
	n := len(arr)
	target := make([]int, n)
	copy(target, arr)
	sort.Ints(target)
	diff := 0 // 不同
	for i := 0; i < n; i++ {
		diff = diff + arr[i] - target[i]
		if diff == 0 { // 累计次数抵消后为0，次数+1
			res++
		}
	}
	return res
}
