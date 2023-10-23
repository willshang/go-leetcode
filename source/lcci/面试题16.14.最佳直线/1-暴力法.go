package main

func main() {

}

// 程序员面试金典16.14_最佳直线
func bestLine(points [][]int) []int {
	res := []int{0, 1}
	maxCount := 0
	n := len(points)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			count := 2
			x1 := points[i][0] - points[j][0]
			y1 := points[i][1] - points[j][1]
			for k := j + 1; k < n; k++ {
				x2 := points[i][0] - points[k][0]
				y2 := points[i][1] - points[k][1]
				if x1*y2 == x2*y1 { // 斜率相同+1
					count++
				}
			}
			if count > maxCount {
				maxCount = count
				res[0] = i
				res[1] = j
			}
		}
	}
	return res
}
