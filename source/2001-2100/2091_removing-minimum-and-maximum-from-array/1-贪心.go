package main

func main() {

}

// leetcode2091_从数组中移除最大值和最小值
func minimumDeletions(nums []int) int {
	minIndex, maxIndex := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
		if nums[i] < nums[minIndex] {
			minIndex = i
		}
	}
	if minIndex > maxIndex { // 保证minIndex < maxIndex
		minIndex, maxIndex = maxIndex, minIndex
	}
	a := maxIndex + 1                        // 第1种情况：全往左移
	b := len(nums) - minIndex                // 第2种情况：全往右移
	c := minIndex + 1 + len(nums) - maxIndex // 第3种情况：往2边移动
	return min(min(a, b), c)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
