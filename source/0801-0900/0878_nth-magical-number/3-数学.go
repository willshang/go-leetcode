package main

func main() {

}

var mod = 1000000007

func nthMagicalNumber(n int, a int, b int) int {
	ab := lcm(a, b)
	// 设 ab为A，B的最小公倍数，如果 x(x<=ab)是神奇数字，那么x+ab也是神奇数字
	m := ab/a + ab/b - 1 // 有m个神奇数字小于ab
	// n = m*q + r
	q := n / m
	r := n % m
	res := q * ab % mod
	if r == 0 {
		return res
	}
	arr := []int{a, b}
	for i := 0; i < r-1; i++ { // 计算剩下的r-1个神奇数字
		if arr[0] <= arr[1] {
			arr[0] = arr[0] + a
		} else {
			arr[1] = arr[1] + b
		}
	}
	res = res + min(arr[0], arr[1])
	return res % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
