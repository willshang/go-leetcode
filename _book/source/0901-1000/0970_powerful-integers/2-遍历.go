package main

import "fmt"

func main() {
	fmt.Println(powerfulIntegers(2, 3, 1000000))
}

func powerfulIntegers(x int, y int, bound int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	if bound < 2 {
		return res
	}
	xArr := make([]int, 0)
	yArr := make([]int, 0)
	for i := 1; i < bound; i = i * x {
		xArr = append(xArr, i)
		if x == 1 {
			break
		}
	}
	for i := 1; i < bound; i = i * y {
		yArr = append(yArr, i)
		if y == 1 {
			break
		}
	}
	for i := 0; i < len(xArr); i++ {
		for j := 0; j < len(yArr); j++ {
			if xArr[i]+yArr[j] <= bound && m[xArr[i]+yArr[j]] == 0 {
				res = append(res, xArr[i]+yArr[j])
				m[xArr[i]+yArr[j]] = 1
			}
		}
	}
	return res
}
