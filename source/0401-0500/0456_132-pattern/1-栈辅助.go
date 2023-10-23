package main

func main() {

}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	minArr := make([]int, len(nums)) // minArr[i]是[0,i]中的最小值
	minArr[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		minArr[i] = min(nums[i], minArr[i-1])
	}
	stack := make([]int, 0)
	for j := len(nums) - 1; j >= 0; j-- {
		// a[i]<a[k]<a[j] => min[j]<a[k]<a[j]
		if nums[j] > minArr[j] {
			for len(stack) > 0 && stack[len(stack)-1] <= minArr[j] {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] < nums[j] {
				return true
			}
			stack = append(stack, nums[j])
		}
	}
	return false
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
