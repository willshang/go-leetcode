package main

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))
}

// leetcode991_坏了的计算器
func brokenCalc(X int, Y int) int {
	if X > Y {
		return X - Y
	}
	res := 0
	for X < Y {
		res++
		if Y%2 == 1 {
			Y++
		} else {
			Y = Y / 2
		}
	}
	return res + X - Y
}
