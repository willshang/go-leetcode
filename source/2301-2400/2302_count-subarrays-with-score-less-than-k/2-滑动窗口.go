package main

func main() {

}

// leetcode2302_统计得分小于K的子数组数目
func countSubarrays(nums []int, k int64) int64 {
	var res, sum int64
	var left int
	for right := 0; right < len(nums); right++ {
		sum = sum + int64(nums[right])
		for sum*int64(right-left+1) >= k {
			sum = sum - int64(nums[left])
			left++
		}
		res = res + int64(right-left+1)
	}
	return res
}
