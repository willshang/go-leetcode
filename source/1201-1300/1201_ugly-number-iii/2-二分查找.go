package main

import "fmt"

func main() {
	fmt.Println(nthUglyNumber(1000000000, 2, 217983653, 336916467))
}

// leetcode1201_丑数III
func nthUglyNumber(n int, a int, b int, c int) int {
	ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
	abc := lcm(ab, c)
	left := 0
	right := 2000000000
	for left < right {
		mid := left + (right-left)/2
		total := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if total >= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 求最小公倍数
func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

// 求最大公约数
func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(a%b, b)
}
