package main

import "math"

func main() {

}

func maximumGroups(grades []int) int {
	n := len(grades)
	// k*(k+1)/2 <= n
	// k^2 + k - 2*n <= 0
	ans := math.Sqrt(0.25+2*float64(n)) - 0.5
	// 求根公式
	// ans = (math.Sqrt(1+8*float64(n)) - 1) / 2
	return int(math.Floor(ans)) // 向下取整
}
