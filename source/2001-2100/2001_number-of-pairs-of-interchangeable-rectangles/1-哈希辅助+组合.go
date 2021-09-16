package main

import "fmt"

func main() {

}

// leetcode2001_可互换矩形的组数
func interchangeableRectangles(rectangles [][]int) int64 {
	m := make(map[string]int64)
	res := int64(0)
	for i := 0; i < len(rectangles); i++ {
		a, b := rectangles[i][0], rectangles[i][1]
		c := gcd(a, b)
		m[fmt.Sprintf("%d,%d", a/c, b/c)]++
	}
	for _, v := range m {
		res = res + (v-1)*v/2
	}
	return res
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
