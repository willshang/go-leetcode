package main

import "fmt"

func main() {
	fmt.Println(waysToSplit([]int{1, 1, 1}))
}

// leetcode1712_将数组分成三个子数组的方案数
func waysToSplit(nums []int) int {
	res := 0
	n := len(nums)
	arr := make([]int, n+1)
	for i := 0; i < n; i++ {
		arr[i+1] = arr[i] + nums[i]
	}
	left, right := 2, 2
	for i := 1; i <= n-1; i++ {
		first := arr[i] // 第1个数
		left = max(left, i+1)
		right = max(right, i+1)
		for left <= n-1 && arr[left]-first < first { // 找到第一个满足的中间数组左边坐标
			left++
		}
		for right <= n-1 && arr[right]-first <= arr[n]-arr[right] {
			right++
		}
		if left <= right {
			res = (res + right - left) % 1000000007
		}
	}
	return res % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
