package main

func main() {

}

// leetcode2025_分割数组的最多方案数
func waysToPartition(nums []int, k int) int {
	res := 0
	n := len(nums)
	prev := make(map[int]int) // 前缀和
	arr := make([]int, n)
	arr[0] = nums[0]
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] + nums[i]
		prev[arr[i-1]]++
	}
	sum := arr[n-1]
	if sum%2 == 0 {
		res = prev[sum/2] // 不改变的结果
	}
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		diff := k - nums[i]
		if (sum+diff)%2 == 0 {
			res = max(res, prev[(sum-diff)/2]+m[(sum+diff)/2])
		}
		m[arr[i]]++    // 左侧+1
		prev[arr[i]]-- // 右侧-1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
