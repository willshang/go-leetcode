package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

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
	str := strconv.Itoa(n)
	for _, v := range str {
		if v == '0' || int32(n)%(v-'0') != 0 {
			return false
		}
	}
	return true
}
