package main

import "math"

func main() {

}

// leetcode1761_一个图中连通三元组的最小度数
func minTrioDegree(n int, edges [][]int) int {
	arr := make([]int, 401)
	m := make(map[int]bool)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a]++
		arr[b]++
		m[getValue(a, b)] = true
	}
	res := math.MaxInt32
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if m[getValue(i, j)] == true {
				for k := j + 1; k <= n; k++ {
					if m[getValue(i, k)] == true && m[getValue(j, k)] == true {
						if arr[i]+arr[j]+arr[k]-6 < res {
							res = arr[i] + arr[j] + arr[k] - 6
						}
					}
				}
			}
		}
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func getValue(a, b int) int {
	if a > b {
		a, b = b, a
	}
	return a*1000 + b
}
