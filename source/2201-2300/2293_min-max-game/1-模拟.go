package main

func main() {

}

func minMaxGame(nums []int) int {
	for len(nums) > 1 {
		temp := make([]int, len(nums)/2)
		for i := 0; i < len(nums)/2; i++ {
			if i%2 == 0 {
				temp[i] = min(nums[i*2], nums[i*2+1])
			} else {
				temp[i] = max(nums[i*2], nums[i*2+1])
			}
		}
		nums = temp
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
