package main

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
}

// leetcode1232_缀点成线
// k=y/x k1=y1/x1 => xy1=x1y
func checkStraightLine(coordinates [][]int) bool {
	x, y := coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]
	for i := 2; i < len(coordinates); i++ {
		x1, y1 := coordinates[i][0]-coordinates[i-1][0], coordinates[i][1]-coordinates[i-1][1]
		if x*y1 != x1*y {
			return false
		}
	}
	return true
}
