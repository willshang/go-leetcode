package main

import (
	"fmt"
)

func main() {
	fmt.Println(singleNumbers([]int{4, 1, 4, 6}))
	fmt.Println(10 & (-10))
}

/*
a = 10
1.a   00001010
2.取反 11110101
3.取反+1 11110110
4. a & (-a)
00001010
11110110
00000010
*/
func singleNumbers(nums []int) []int {
	res := make([]int, 2)
	temp := 0
	for i := 0; i < len(nums); i++ {
		temp = temp ^ nums[i]
	}
	last := temp & (-temp)
	for i := 0; i < len(nums); i++ {
		if nums[i]&last == 0 {
			res[0] = res[0] ^ nums[i]
		} else {
			res[1] = res[1] ^ nums[i]
		}
	}
	return res
}
