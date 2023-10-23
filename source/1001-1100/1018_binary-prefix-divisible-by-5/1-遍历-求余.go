package main

import "fmt"

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))
}

// leetcode1018_可被5整除的二进制前缀
func prefixesDivBy5(A []int) []bool {
	res := make([]bool, len(A))
	temp := 0
	for i := 0; i < len(A); i++ {
		temp = (temp*2 + A[i]) % 5
		if temp == 0 {
			res[i] = true
		}
	}
	return res
}
