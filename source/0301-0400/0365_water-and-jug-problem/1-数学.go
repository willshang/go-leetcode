package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))
}

// leetcode365_水壶问题
// ax+by=z
func canMeasureWater(x int, y int, z int) bool {
	if x > y {
		x, y = y, x
	}
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x, y) == 0
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
