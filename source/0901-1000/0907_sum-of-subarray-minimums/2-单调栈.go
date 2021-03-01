package main

func main() {

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
		res = (res + total) % 1000000007
	}
	return res % 1000000007
}
