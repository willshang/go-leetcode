package main

import "fmt"

func main() {
	fmt.Println(simplifiedFractions(6))
}

// leetcode1447_最简分数
func simplifiedFractions(n int) []string {
	res := make([]string, 0)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				res = append(res, fmt.Sprintf("%d/%d", j, i))
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}
