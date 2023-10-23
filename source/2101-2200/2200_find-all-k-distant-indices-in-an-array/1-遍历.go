package main

func main() {

}

// leetcode2200_找出数组中的所有K近邻下标
func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if nums[j] == key && abs(i-j) <= k {
				res = append(res, i)
				break
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
