package main

import "fmt"

func main() {
	fmt.Println(maxPoints([][]int{{1, 1}, {2, 2}, {3, 3}}))
}

func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}
	res := 2
	m := make(map[[3]int]map[int]bool)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// AX+BY+C=0
			A := points[j][1] - points[i][1]
			B := points[i][0] - points[j][0]
			C := points[i][1]*points[j][0] - points[i][0]*points[j][1]
			com := gcd(gcd(A, B), C)
			A, B, C = A/com, B/com, C/com
			node := [3]int{A, B, C}
			if m[node] == nil {
				m[node] = make(map[int]bool)
			}
			m[node][i] = true
			m[node][j] = true
			if len(m[node]) > res {
				res = len(m[node])
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
