package main

import "fmt"

func main() {
	fmt.Println(totalHammingDistance([]int{4, 14, 2}))
}

// leetcode477_汉明距离总和
func totalHammingDistance(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		total := 0
		for j := 0; j < len(nums); j++ {
			total = total + (nums[j]>>i)&1
		}
		// 汉明距离：两个数字的二进制数对应位不同的数量
		// 该位1的数量*该位0的数量
		res = res + total*(len(nums)-total)
	}
	return res
}
