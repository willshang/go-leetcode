package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
}

var res [][]int
var arr []int

func maxCoins(nums []int) int {
	n := len(nums)
	arr = make([]int, n+2)
	arr[0], arr[len(arr)-1] = 1, 1
	for i := 1; i <= n; i++ {
		arr[i] = nums[i-1]
	}
	res = make([][]int, n+2)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n+2)
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = -1
		}
	}
	return dfs(0, n+1)
}

// 将开区间(i,j)内的位置全部填满气球能够得到的最多硬币数
func dfs(left, right int) int {
	// 不满足3个
	if left+1 >= right {
		return 0
	}
	if res[left][right] != -1 {
		return res[left][right]
	}
	for i := left + 1; i < right; i++ {
		// 填充第i位，两边是arr[left],arr[right]
		sum := arr[left] * arr[i] * arr[right]
		sum = sum + dfs(left, i) + dfs(i, right)
		res[left][right] = max(res[left][right], sum)
	}
	return res[left][right]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
