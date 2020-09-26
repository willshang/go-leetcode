package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
}

func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		res[i] = res[i&(i-1)] + 1
	}
	return res
}
