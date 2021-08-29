package main

import "sort"

func main() {

}

// leetcode1979_找出数组的最大公约数
func findGCD(nums []int) int {
	sort.Ints(nums)
	a, b := nums[0], nums[len(nums)-1]
	return gcd(a, b)
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
