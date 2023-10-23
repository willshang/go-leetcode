package main

func main() {

}

// leetcode2176_统计数组中相等且可以被整除的数对
func countPairs(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				res++
			}
		}
	}
	return res
}
