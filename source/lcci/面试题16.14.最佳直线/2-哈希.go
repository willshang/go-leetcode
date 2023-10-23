package main

import "fmt"

func main() {

	fmt.Println(bestLine([][]int{{-1, 2}, {-3, 4}}))
}

func bestLine(points [][]int) []int {
	res := []int{0, 1}
	maxCount := 0
	n := len(points)
	m := make(map[[3]int]int)
	mToKey := make(map[[3]int][]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			// AX+BY+C=0
			A := points[j][1] - points[i][1]
			B := points[i][0] - points[j][0]
			C := points[i][1]*points[j][0] - points[i][0]*points[j][1]
			com := gcd(gcd(A, B), C)
			A, B, C = A/com, B/com, C/com
			node := [3]int{A, B, C}
			if m[node] == 0 {
				mToKey[node] = []int{i, j}
			}
			m[node]++
			if m[node] > maxCount {
				maxCount = m[node]
				res = mToKey[node]
			} else if m[node] == maxCount {
				if mToKey[node][0] < res[0] ||
					(mToKey[node][0] == res[0] && mToKey[node][1] < res[1]) {
					res = mToKey[node]
				}
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
