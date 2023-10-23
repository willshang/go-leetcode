package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{1, 2, 1, 3, 2, 5}))
}

func singleNumber(nums []int) []int {
	a := 0
	for i := 0; i < len(nums); i++ {
		a = a ^ nums[i]
	}
	b := 1
	for a&1 == 0 {
		a = a >> 1
		b = b << 1
	}
	res := []int{0, 0}
	for i := 0; i < len(nums); i++ {
		if nums[i]&b == 0 {
			res[0] = res[0] ^ nums[i]
		} else {
			res[1] = res[1] ^ nums[i]
		}
	}
	return res
}
