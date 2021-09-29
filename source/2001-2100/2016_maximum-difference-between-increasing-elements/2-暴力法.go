package main

func main() {

}

func maximumDifference(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				res = max(res, nums[j]-nums[i])
			}
		}
	}
	if res == 0 {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
