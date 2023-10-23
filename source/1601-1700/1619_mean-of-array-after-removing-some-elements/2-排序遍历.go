package main

import "sort"

func main() {

}

// leetcode1619_删除某些元素后的数组均值
func trimMean(arr []int) float64 {
	n := len(arr)
	count := n / 20
	sort.Ints(arr)
	sum := 0
	for i := count; i < n-count; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(n-2*count)
}
