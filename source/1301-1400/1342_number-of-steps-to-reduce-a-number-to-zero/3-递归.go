package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
	fmt.Println(numberOfSteps(8))
}

// leetcode1342_将数字变成0的操作次数
func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	} else if num%2 == 1 {
		return 1 + numberOfSteps(num-1)
	}
	return 1 + numberOfSteps(num/2)
}
