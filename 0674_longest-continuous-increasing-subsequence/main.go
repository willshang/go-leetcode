package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 4, 7}
	fmt.Println(findLengthOfLCIS(arr))
}
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 1
	i, j := 0, 1
	for j < len(nums) {
		for j < len(nums) && nums[j-1] < nums[j] {
			j++
		}
		if res < j-i {
			res = j - i
		}
		i = j
		j++
	}
	return res
}
