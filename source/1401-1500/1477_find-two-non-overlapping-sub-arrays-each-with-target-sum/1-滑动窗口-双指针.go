package main

import "math"

func main() {

}

// leetcode1477_找两个和为目标值且不重叠的子数组
func minSumOfLengths(arr []int, target int) int {
	res := math.MaxInt32
	n := len(arr)
	sum := 0
	left := 0
	temp := make([]int, n) // 保存每个位置之前的最小值
	for i := 0; i < n; i++ {
		temp[i] = math.MaxInt32 // 默认为最大值
	}
	for right := 0; right < n; right++ {
		sum = sum + arr[right]
		for sum > target {
			sum = sum - arr[left]
			left++
		}
		if right >= 1 {
			temp[right] = temp[right-1] // 默认同之前最小值
		}
		if sum == target { // 找到目标值
			temp[right] = min(temp[right], right-left+1) // 更新最小值
			if left >= 1 && temp[left-1] != math.MaxInt32 {
				res = min(res, temp[left-1]+right-left+1) // 取left之前(即left-1)的最小值，加上当前长度
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
