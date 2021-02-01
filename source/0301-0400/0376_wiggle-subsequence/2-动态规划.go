package main

func main() {

}

// leetcode376_摆动序列
func wiggleMaxLength(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := 1   // 以前i个元素中的某一个为结尾的最长的上升摆动序列的长度
	down := 1 // 以前i个元素中的某一个为结尾的最长的下降摆动序列的长度。
	for i := 1; i < n; i++ {
		if nums[i-1] < nums[i] { // 递增
			up = down + 1
		} else if nums[i-1] > nums[i] { // 递减

			down = up + 1
		}
	}
	return max(up, down)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
