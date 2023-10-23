package main

func main() {

}

// leetcode2090_半径为k的子数组平均值
func getAverages(nums []int, k int) []int {
	n := len(nums)
	res := make([]int, n)
	left := 0
	sum := 0
	for right := 0; right < n; right++ {
		res[right] = -1
		sum = sum + nums[right]
		if right >= 2*k {
			res[right-k] = sum / (2*k + 1)
			sum = sum - nums[left]
			left++
		}
	}
	return res
}
