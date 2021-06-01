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
	for i := a; i < n-b; i++ { // 枚举L以i结尾+右边（固定长度为b）
		L = max(L, nums[i]-nums[i-a])   // L为结尾的数组中的L最大值
		temp := L + nums[i+b] - nums[i] // L+M
		res = max(res, temp)
	}
	for i := b; i < n-a; i++ { // 枚举M以i结尾+右边（固定长度为a）
		M = max(M, nums[i]-nums[i-b])   // M为i结尾的数组中的M最大值
		temp := nums[i+a] - nums[i] + M //  M+L
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
