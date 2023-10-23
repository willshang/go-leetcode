package main

func main() {

}

func nthUglyNumber(n int, a int, b int, c int) int {
	ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
	abc := lcm(ab, c)
	left := 1
	right := 2000000000
	for left <= right {
		mid := left + (right-left)/2
		total := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
		if total == n {
			if mid%a == 0 || mid%b == 0 || mid%c == 0 {
				return mid
			}
			right = mid - 1
		} else if total < n {
			left = mid + 1
		} else {
			right = mid - 1
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
