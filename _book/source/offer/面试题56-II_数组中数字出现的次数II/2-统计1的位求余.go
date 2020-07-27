package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{3, 4, 3, 3}))
}

// 剑指offer_面试题56-II.数组中数字出现的次数II
func singleNumber(nums []int) int {
	arr := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		value := nums[i]
		for j := 31; j >= 0; j-- {
			arr[j] = arr[j] + value&1
			value = value / 2
		}
	}
	res := 0
	for i := 0; i < 32; i++ {
		res = res << 1
		res = res + arr[i]%3
	}
	return res
}
