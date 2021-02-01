package main

func main() {

}

func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := make([]int, n)   // 以前i个元素中的某一个为结尾的最长的上升摆动序列的长度
	down := make([]int, n) // 以前i个元素中的某一个为结尾的最长的下降摆动序列的长度。
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] { // 递增
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i-1] > nums[i] { // 递减
			up[i] = up[i-1]
			down[i] = max(up[i-1]+1, down[i-1])
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
