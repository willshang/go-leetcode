package main

func main() {

}

// leetcode2161_根据给定数字划分数组
func pivotArray(nums []int, pivot int) []int {
	n := len(nums)
	a, b, c := make([]int, 0), make([]int, 0), make([]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] < pivot {
			a = append(a, nums[i])
		} else if nums[i] == pivot {
			b = append(b, nums[i])
		} else {
			c = append(c, nums[i])
		}
	}
	return append(a, append(b, c...)...)
}
