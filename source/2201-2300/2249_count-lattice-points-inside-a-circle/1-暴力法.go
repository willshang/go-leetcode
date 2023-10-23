package main

func main() {

}

// leetcode2249_统计圆内格点数目
func countLatticePoints(circles [][]int) int {
	res := 0
	for i := 0; i <= 200; i++ {
		for j := 0; j <= 200; j++ {
			for k := 0; k < len(circles); k++ {
				x, y, r := circles[k][0], circles[k][1], circles[k][2]
				if (i-x)*(i-x)+(j-y)*(j-y) <= r*r {
					res++
					break
				}
			}
		}
	}
	return res
}
