package main

import "math"

func main() {

}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr[i][j] = math.MaxInt32
		}
		arr[i][i] = 0
	}
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1], edges[i][2] // a<=>b
		arr[a][b], arr[b][a] = c, c
	}
	for k := 0; k < n; k++ { // floyd
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][k] != math.MaxInt32 && arr[k][j] != math.MaxInt32 &&
					arr[i][j] >= arr[i][k]+arr[k][j] {
					arr[i][j] = arr[i][k] + arr[k][j]
				}
			}
		}
	}
	res := -1
	minCount := math.MaxInt32
	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < n; j++ {
			if arr[i][j] <= distanceThreshold && i != j {
				count++
			}
		}
		if count <= minCount {
			minCount = count
			res = i // 等于的时候更新为较大的编号
		}
	}
	return res
}
