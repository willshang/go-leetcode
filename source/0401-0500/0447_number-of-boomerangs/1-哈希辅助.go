package main

import "fmt"

func main() {
	points := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
	}

	fmt.Println(numberOfBoomerangs(points))
}

// leetcode447_回旋镖的数量
func numberOfBoomerangs(points [][]int) int {
	res := 0
	size := len(points)
	if size < 3 {
		return 0
	}
	for i := 0; i < size; i++ {
		dMap := make(map[int]int, size)
		for j := 0; j < size; j++ {
			if i == j {
				continue
			}
			d := dSquare(points[i], points[j])
			if _, ok := dMap[d]; ok {
				dMap[d]++
			} else {
				dMap[d] = 1
			}
		}
		// 相同距离的v个点，总共有 v*(v-1)种排列
		for _, v := range dMap {
			res = res + v*(v-1)
		}
	}
	return res
}

func dSquare(p1, p2 []int) int {
	x := p2[0] - p1[0]
	y := p2[1] - p1[1]
	return x*x + y*y
}
