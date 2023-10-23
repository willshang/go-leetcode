package main

func main() {

}

// leetcode2202_K次操作后最大化顶端元素
func maximumTop(nums []int, k int) int {
	n := len(nums)
	if n == 1 && k%2 == 1 { // 特殊情况：操作1次删除后栈为空
		return -1
	}
	res := 0
	// 2种情况：
	// 1、k-1中的最大值，最后1次选放回最大值（不会选中第k个数）
	// 2、k+1个位置，删除k个数，剩下的就是栈顶
	for i := 1; i <= n; i++ {
		if i <= k+1 && i != k {
			res = max(res, nums[i-1])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
