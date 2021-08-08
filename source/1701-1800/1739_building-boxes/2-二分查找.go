package main

import "fmt"

func main() {
	fmt.Println(minimumBoxes(37910271))
}

func minimumBoxes(n int) int {
	res := 0
	left, right := 1, 3000
	for left < right {
		mid := left + (right-left)/2
		if mid*(mid+1)*(mid+2)/6 < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	left = left - 1
	res = (1 + left) * left / 2
	n = n - left*(left+1)*(left+2)/6
	left, right = 0, n
	for left < right {
		mid := left + (right-left)/2
		target := mid * (mid + 1) / 2
		if target < n {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return res + left
}
