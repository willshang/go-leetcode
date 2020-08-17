package main

import (
	"fmt"
)

func main() {
	fmt.Println(string(findKthBit(4, 12)))
}

func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	flag := false
	mid := 1 << (n - 1)
	for k > 1 {
		if k == mid {
			if flag == true {
				return '0'
			}
			return '1'
		} else if k > mid {
			if flag == true {
				flag = false
			} else {
				flag = true
			}
			k = (mid << 1) - k
		}
		mid = mid >> 1
	}
	if flag == true {
		return '1'
	}
	return '0'
}
