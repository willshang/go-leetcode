package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Print(binarySearch([]int{1, 2, 3, 3, 4, 4, 10, 10, 12}, 5))
}

var mod = 1000000007

func minWastedSpace(packages []int, boxes [][]int) int {
	sort.Ints(packages)
	n := len(packages)
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
			right := binarySearch(packages, temp[j]) // 选择当前箱子能装下的位置
			sum = sum + (right-left)*temp[j]         // 累加：个数*箱子大小
			left = right
		}
		res = min(res, sum)
	}
	if res == math.MaxInt64 {
		return -1
	}
	for i := 0; i < n; i++ {
		res = res - packages[i]
	}
	return res % mod
}

// 找到第一个大于target的位置
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
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
