package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{
		{44, 31, 4},
		{47, 27, 27},
		{7, 13, 0},
		{13, 21, 20},
		{50, 34, 18},
		{47, 44, 28},
	}
	fmt.Println(bestCoordinate(arr, 13))
}

// leetcode1620_网络信号最好的坐标
func bestCoordinate(towers [][]int, radius int) []int {
	res := []int{0, 0}
	maxValue := 0
	for i := 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			sum := 0
			for k := 0; k < len(towers); k++ {
				a, b, c := towers[k][0], towers[k][1], towers[k][2]
				d := (a-i)*(a-i) + (b-j)*(b-j)
				if d <= radius*radius {
					value := float64(c) / (1 + math.Sqrt(float64(d)))
					sum = sum + int(math.Floor(value))
				}
			}
			if sum > maxValue {
				maxValue = sum
				res = []int{i, j}
			}
		}
	}
	return res
}
