package main

func main() {

}

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	res := 0
	n := len(nums)
	a, b := firstLen, secondLen
	for i := 1; i < n; i++ {
		nums[i] = nums[i] + nums[i-1] // 前缀和
	}
	res = nums[a+b-1]
	L, M := nums[a-1], nums[b-1]
	for i := a + b; i < n; i++ { // 枚举以i结尾的数组：分为左边（不固定）+右边（固定长度为a/b）
		L = max(L, nums[i-b]-nums[i-b-a])                     // L为i-b结尾的数组中的L最大值
		M = max(M, nums[i-a]-nums[i-a-b])                     // M为i-a结尾的数组中的M最大值
		temp := max(L+nums[i]-nums[i-b], nums[i]-nums[i-a]+M) // L+M 和 M+L的较大者
		res = max(res, temp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
