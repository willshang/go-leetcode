package main

import "fmt"

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

// leetcode728_自除数
func selfDividingNumbers(left int, right int) []int {
	res := make([]int, 0)
	for i := left; i <= right; i++ {
		if isSelfDividing(i) {
			res = append(res, i)
		}
	}
	return res
}

func isSelfDividing(n int) bool {
	temp := n
	for temp > 0 {
		d := temp % 10
		temp = temp / 10
		if d == 0 || n%d != 0 {
			return false
		}
	}
	return true
}
