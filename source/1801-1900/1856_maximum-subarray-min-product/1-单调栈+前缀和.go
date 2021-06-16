package main

func main() {

}

var mod = 1000000007

func maxSumMinProduct(nums []int) int {
	res := 0
	n := len(nums)
	arr := make([]int, n+1) // 前缀和
	for i := 1; i <= n; i++ {
		arr[i] = arr[i-1] + nums[i-1]
	}
	left := make([]int, n)  // 左侧最近的严格小于nums[i]的元素下标
	right := make([]int, n) // 右侧最近的小于等于nums[i]的元素下标
	for i := 0; i < n; i++ {
		left[i] = 0      // 默认是最左边
		right[i] = n - 1 // 默认是最右边
	}
	stack := make([]int, 0) // 单调递减栈
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			right[stack[len(stack)-1]] = i - 1
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1] + 1
		}
		stack = append(stack, i)
	}
	for i := 0; i < n; i++ {
		target := (arr[right[i]+1] - arr[left[i]]) * nums[i]
		res = max(res, target)
	}
	return res % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
