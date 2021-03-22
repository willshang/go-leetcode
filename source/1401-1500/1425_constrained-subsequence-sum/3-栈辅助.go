package main

func main() {

}

// leetcode1425_带限制的子序列和
func constrainedSubsetSum(nums []int, k int) int {
	n := len(nums)
	if k > n {
		k = n
	}
	temp := nums[0]
	res := nums[0]
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{0, nums[0]})
	for i := 1; i < len(nums); i++ {
		if stack[0][0] < i-k {
			stack = stack[1:]
		}
		if stack[0][1] < 0 {
			temp = nums[i]
		} else {
			temp = stack[0][1] + nums[i]
		}
		for len(stack) > 0 && stack[len(stack)-1][1] < temp {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, [2]int{i, temp})
		if temp > res {
			res = temp
		}
	}
	return res
}
