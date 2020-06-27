package main

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}))
}

// 鞋带公式
// S=|(x1 * y2 + x2 * y3 + x3 * y1 - y1 * x2 - y2 * x3 - y3 * x1)|/2
// S==0组成一条直线
func checkStraightLine(coordinates [][]int) bool {
	for i := 2; i < len(coordinates); i++ {
		x1 := coordinates[i-2][0]*coordinates[i-1][1] +
			coordinates[i-1][0]*coordinates[i][1] + coordinates[i][0]*coordinates[i-2][1]
		x2 := coordinates[i-2][1]*coordinates[i-1][0] +
			coordinates[i-1][1]*coordinates[i][0] + coordinates[i][1]*coordinates[i-2][0]
		if x1 != x2 {
			return false
		}
	}
	return true
}
