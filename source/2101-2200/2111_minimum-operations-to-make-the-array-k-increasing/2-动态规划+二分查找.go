package main

func main() {

}

// leetcode2111_使数组K递增的最少操作次数
func kIncreasing(arr []int, k int) int {
	res := 0
	for i := 0; i < k; i++ {
		dp := make([]int, 0)
		count := 0
		for j := i; j < len(arr); j = j + k {
			count++
			index := upperBound(dp, arr[j])
			if index == len(dp) {
				dp = append(dp, arr[j])
			} else {
				dp[index] = arr[j]
			}
		}
		res = res + count - len(dp)
	}
	return res
}

// 返回第一个大于等于target的位置
func upperBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1 // 收缩左边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
