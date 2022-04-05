package main

func main() {

}

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	res := make([]int, 0)
	right := 0
	for j := 0; j < n; j++ {
		if nums[j] == key {
			left := max(right, j-k)   // 往前k or 已经访问过的
			right = min(n-1, j+k) + 1 // 更新
			for i := left; i < right; i++ {
				res = append(res, i)
			}
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
