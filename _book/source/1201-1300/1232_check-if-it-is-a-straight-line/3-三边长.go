package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println(checkStraightLine([][]int{{2, 1}, {4, 2}, {6, 3}}))
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
	//fmt.Println(checkStraightLine([][]int{{0, 1}, {-6, -17}, {3, 10}, {-7, -20}, {1, 4}, {4, 13}, {9, 28}, {7, 22}}))
}

func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		side1 := side(coordinates[i], coordinates[i-1])
		side2 := side(coordinates[i-1], coordinates[i-2])
		side3 := side(coordinates[i], coordinates[i-2])
		arr := []float64{side1, side2, side3}
		sort.Float64s(arr)
		if arr[2]-arr[1]-arr[0] > 0.00000005 || arr[2]-arr[1]-arr[0] < -0.00000005 {
			return false
		}
	}
	return true
}

func side(arr1, arr2 []int) float64 {
	res := (arr1[0]-arr2[0])*(arr1[0]-arr2[0]) +
		(arr1[1]-arr2[1])*(arr1[1]-arr2[1])
	return math.Sqrt(float64(res))
}
