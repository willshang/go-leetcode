package main

func main() {

}

// leetcode2006_差的绝对值为K的数对数目
func countKDifference(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]-nums[j] == k || nums[j]-nums[i] == k {
				res++
			}
		}
	}
	return res
}
