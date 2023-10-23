package main

import "sort"

func main() {

}

func findFinalValue(nums []int, original int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == original {
			original = original * 2
		}
	}
	return original
}
