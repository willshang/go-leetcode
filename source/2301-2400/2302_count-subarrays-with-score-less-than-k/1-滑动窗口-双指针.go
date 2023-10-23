package main

func main() {

}

func countSubarrays(nums []int, k int64) int64 {
	var res, sum, count int64
	n := len(nums)
	left := -1
	for right := 0; right < n; right++ {
		sum = sum + int64(nums[right])
		count++
		for left < n && sum*count >= k {
			left++
			sum = sum - int64(nums[left])
			count--
		}
		res = res + count
	}
	return res
}
