package main

import "math"

func main() {

}

// leetcode878_第N个神奇数字
var mod = 1000000007

func nthMagicalNumber(n int, a int, b int) int {
	ab := lcm(a, b)
	left := 0
	right := int(math.Pow10(15))
	for left < right {
		mid := left + (right-left)/2
		total := mid/a + mid/b - mid/ab
		if total >= n {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left % mod
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
