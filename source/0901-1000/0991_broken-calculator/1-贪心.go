package main

import "fmt"

func main() {
	fmt.Println(brokenCalc(2, 3))
}

func brokenCalc(X int, Y int) int {
	if X > Y {
		return X - Y
	}
	res := 0
	for X < Y {
		if Y%2 == 0 {
			Y = Y / 2
			res++
		} else {
			Y = (Y + 1) / 2
			res = res + 2
		}
	}
	return res + X - Y
}
