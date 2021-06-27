package main

import "math"

func main() {

}

// leetcode1906_查询差绝对值的最小值
func minDifference(nums []int, queries [][]int) []int {
	n := len(nums)
	arr := make([][101]int, n+1)
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1]
		arr[i][nums[i-1]]++
	}
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = -1
		left, right := queries[i][0], queries[i][1]
		prev, minValue := 0, math.MaxInt32
		for j := 1; j <= 100; j++ { // 枚举每个数
			if arr[right+1][j] != arr[left][j] { // 当j出现
				if prev != 0 {
					if j-prev < minValue { // 保存较小的差值
						minValue = j - prev
					}
				}
				prev = j // 更新上一个出现的数
			}
		}
		if minValue != math.MaxInt32 {
			res[i] = minValue
		}
	}
	return res
}
