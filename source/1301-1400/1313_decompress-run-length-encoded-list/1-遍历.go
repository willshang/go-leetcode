package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
}

// leetcode1313_解压缩编码列表
func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums); i = i + 2 {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
	}
	return res
}
