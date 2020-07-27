package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}}))
	//fmt.Println(isBoomerang([][]int{{0, 0}, {1, 0}, {2, 2}}))
	fmt.Println(isBoomerang([][]int{{0, 0}, {1, 2}, {0, 1}}))
}

func isBoomerang(points [][]int) bool {
	side1 := side(points[0], points[1])
	side2 := side(points[1], points[2])
	side3 := side(points[0], points[2])
	return side1+side2 > side3 &&
		side2+side3 > side1 &&
		side1+side3 > side2
}

func side(arr1, arr2 []int) float64 {
	res := (arr1[0]-arr2[0])*(arr1[0]-arr2[0]) +
		(arr1[1]-arr2[1])*(arr1[1]-arr2[1])
	return math.Sqrt(float64(res))
}
