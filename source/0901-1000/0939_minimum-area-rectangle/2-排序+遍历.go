package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

}

func minAreaRect(points [][]int) int {
	res := math.MaxInt32
	m := make(map[int][]int)
	n := len(points)
	for i := 0; i < n; i++ {
		a, b := points[i][0], points[i][1]
		m[a] = append(m[a], b) // 列分组
	}
	rows := make([]int, 0)
	for k := range m {
		rows = append(rows, k)
	}
	sort.Ints(rows)
	prevM := make(map[string]int)
	for _, v := range rows { // 按列枚举
		arr := m[v]
		sort.Ints(arr)
		for i := 0; i < len(arr); i++ {
			for j := i + 1; j < len(arr); j++ {
				y1, y2 := arr[i], arr[j]
				target := fmt.Sprintf("%d,%d", y1, y2)
				if index, ok := prevM[target]; ok {
					area := (y2 - y1) * (v - index)
					res = min(res, area)
				}
				prevM[target] = v
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
