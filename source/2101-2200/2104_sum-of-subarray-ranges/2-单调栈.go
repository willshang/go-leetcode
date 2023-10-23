package main

import "fmt"

func main() {
	fmt.Println(subArrayRanges([]int{1, 2, 3}))
}

func subArrayRanges(nums []int) int64 {
	return int64(sumSubarrayMaxs(nums)) - int64(sumSubarrayMins(nums))
}

// leetcode907_子数组的最小值之和
func sumSubarrayMins(arr []int) int {
	res := 0
	stack := make([]int, 0) // 递增栈
	stack = append(stack, -1)
	total := 0
	for i := 0; i < len(arr); i++ {
		for len(stack) > 1 && arr[i] < arr[stack[len(stack)-1]] { // 小于栈顶
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			total = total - arr[prev]*(prev-stack[len(stack)-1])
		}
		stack = append(stack, i)
		total = total + arr[i]*(i-stack[len(stack)-2])
		res = res + total
	}
	return res
}

func sumSubarrayMaxs(arr []int) int {
	res := 0
	stack := make([]int, 0) // 递减栈
	stack = append(stack, -1)
	total := 0
	for i := 0; i < len(arr); i++ {
		for len(stack) > 1 && arr[i] > arr[stack[len(stack)-1]] { // 大于栈顶
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			total = total - arr[prev]*(prev-stack[len(stack)-1])
		}
		stack = append(stack, i)
		total = total + arr[i]*(i-stack[len(stack)-2])
		res = res + total
	}
	return res
}
