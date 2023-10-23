package main

import (
	"math"
	"sort"
)

func main() {

}

var mod = 1000000007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
	arr := make([]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + packages[i-1]
	}
	res := math.MaxInt64
	for i := 0; i < len(boxes); i++ {
		temp := boxes[i]
		sort.Ints(temp)
		if temp[len(temp)-1] < packages[n-1] { // 最大箱子无法装最大包裹
			continue
		}
		sum := 0
		left := 0
		for j := 0; j < len(temp); j++ { // 枚举箱子
			right := binarySearch(packages, temp[j], left) // 选择当前箱子能装下的位置
			sum = sum + (right-left)*temp[j] - (arr[right] - arr[left])
			left = right
		}
		res = min(res, sum)
	}
	if res == math.MaxInt64 {
		return -1
	}
	return res % mod
}

func binarySearch(arr []int, target int, left int) int {
	right := len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
