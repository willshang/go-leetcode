package main

func main() {

}

// leetcode1401_圆和矩形是否有重叠
func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	minValue := 0
	// minValue = (x-x_center)*(x-x_center)+(y-y_center)*(y-y_center)
	if x_center < x1 || x_center > x2 {
		minValue = minValue + min((x1-x_center)*(x1-x_center), (x2-x_center)*(x2-x_center))
	}
	if y_center < y1 || y_center > y2 {
		minValue = minValue + min((y1-y_center)*(y1-y_center), (y2-y_center)*(y2-y_center))
	}
	return minValue <= radius*radius
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
