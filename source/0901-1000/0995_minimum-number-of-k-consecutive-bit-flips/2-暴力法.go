package main

func main() {

}

func minKBitFlips(nums []int, k int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			if i+k > n {
				return -1
			}
			for j := i; j < i+k; j++ {
				nums[j] = 1 - nums[j]
			}
			res++
		}
	}
	return res
}
