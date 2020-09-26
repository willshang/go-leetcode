package main

import "fmt"

func main() {
	fmt.Println(missingTwo([]int{2, 3}))
}

func missingTwo(nums []int) []int {
	n := len(nums) + 2
	sum := (1 + n) * n / 2
	total := 0
	for i := 0; i < len(nums); i++ {
		total = total + nums[i]
	}
	diff := sum - total // a+b
	mid := diff / 2     // (a+b)/2
	tempSum := (1 + mid) * mid / 2
	temp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= mid {
			temp = temp + nums[i]
		}
	}
	a := tempSum - temp
	b := diff - a
	return []int{a, b}
}
