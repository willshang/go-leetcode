package main

func main() {

}

// leetcode2293_极大极小游戏
func minMaxGame(nums []int) int {
	n := len(nums)
	for n > 1 {
		for i := 0; i < n/2; i++ {
			if i%2 == 0 {
				nums[i] = min(nums[i*2], nums[i*2+1])
			} else {
				nums[i] = max(nums[i*2], nums[i*2+1])
			}
		}
		n = n / 2
	}
	return nums[0]
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
