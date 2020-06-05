package main

import "fmt"

func main() {
	fmt.Println(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
}

// leetcode914_卡牌分组
func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	m := make(map[int]int)
	for i := 0; i < len(deck); i++ {
		m[deck[i]]++
	}
	v := m[deck[0]]
	for _, value := range m {
		v = gcd(v, value)
		if v < 2 {
			return false
		}
	}
	return true
}

func gcd(x, y int) int {
	a := x % y
	if a > 0 {
		return gcd(y, a)
	}
	return y
}
