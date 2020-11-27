package main

import "fmt"

func main() {
	fmt.Println(gcd1(5, 10))
	fmt.Println(gcd2(5, 10))
	fmt.Println(gcd3(5, 10))
}

// 辗转相除法
func gcd1(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd1(a%b, b)
}

// 更相减损术
func gcd2(a, b int) int {
	if a == b {
		return a
	}
	if a < b {
		a, b = b, a
	}
	return gcd2(a-b, b)
}

// 更相减损术+移位
func gcd3(a, b int) int {
	if a == b {
		return a
	}
	if a%2 == 0 && b%2 == 0 {
		return gcd3(a>>1, b>>1) << 1
	} else if a%2 == 0 && b%2 == 1 {
		return gcd3(a>>1, b)
	} else if a%2 == 1 && b%2 == 0 {
		return gcd3(a, b>>1)
	}
	if a < b {
		a, b = b, a
	}
	return gcd3(a-b, b)
}
