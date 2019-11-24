package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(nums, 2))
	fmt.Println(nums)
}

/*func removeElement(nums []int, val int) int {
	count := 0
	for i := 0; i < len(nums); i++{
		if nums[i] == val{
			count++
			continue
		}
		nums[i-count] = nums[i]
	}
	return len(nums)-count
}*/

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}
		for j >= 0 && nums[j] == val {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return i
}
