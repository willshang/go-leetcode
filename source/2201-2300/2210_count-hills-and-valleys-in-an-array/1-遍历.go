package main

func main() {

}

func countHillValley(nums []int) int {
	res := 0
	n := len(nums)
	for i := 1; i < n-1; i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		a, b := -1, -1
		for j := i - 1; j >= 0; j-- {
			if nums[i] != nums[j] {
				a = nums[j]
				break
			}
		}
		for j := i + 1; j < n; j++ {
			if nums[i] != nums[j] {
				b = nums[j]
				break
			}
		}
		if a == -1 || b == -1 {
			continue
		}
		if (a < nums[i] && b < nums[i]) || (a > nums[i] && b > nums[i]) {
			res++
		}
	}
	return res
}
