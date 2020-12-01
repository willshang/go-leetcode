package main

func main() {

}

// leetcode1673_找出最具竞争力的子序列
func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0)
	k = len(nums) - k
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		// 栈顶元素打大于后面的元素，摘除栈顶元素（因为前面的更大，需要删除了才能变的最小）
		for len(stack) > 0 && stack[len(stack)-1] > value && k > 0 {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, value)
	}
	stack = stack[:len(stack)-k]
	return stack
}
