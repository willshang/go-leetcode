package main

import "fmt"

func main() {
	fmt.Println(binaryGap(22))
}

// leetcode868_二进制间距
func binaryGap(N int) int {
	res := 0
	count := 0
	for N > 0 {
		if N%2 == 1 {
			if count > res {
				res = count
			}
			count = 1
		} else if count > 0 {
			count++
		}
		N = N / 2
	}
	return res
}
