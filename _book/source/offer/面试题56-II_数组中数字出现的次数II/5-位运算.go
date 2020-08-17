package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{3, 3, 3}))
}

func singleNumber(nums []int) int {
	res, temp := 0, 0
	for i := 0; i < len(nums); i++ {
		res = (res ^ nums[i]) & ^temp
		temp = (temp ^ nums[i]) & ^res
	}
	return res
}
