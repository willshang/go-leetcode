package main

func main() {

}

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	if x1 <= x_center && x_center <= x2 && // 在里面
		y1 <= y_center && y_center <= y2 {
		return true
	} else if x_center > x2 && y1 <= y_center && y_center <= y2 { // 右边
		return x_center-x2 <= radius
	} else if x_center < x1 && y1 <= y_center && y_center <= y2 { // 左边
		return x1-x_center <= radius
	} else if y_center < y1 && x1 <= x_center && x_center <= x2 { // 下边
		return y1-y_center <= radius
	} else if y_center > y2 && x1 <= x_center && x_center <= x2 { // 上边
		return y_center-y2 <= radius
	}
	// 4个顶点判断
	minValue := (x1-x_center)*(x1-x_center) + (y1-y_center)*(y1-y_center)
	minValue = min(minValue, (x1-x_center)*(x1-x_center)+(y2-y_center)*(y2-y_center))
	minValue = min(minValue, (x2-x_center)*(x2-x_center)+(y1-y_center)*(y1-y_center))
	minValue = min(minValue, (x2-x_center)*(x2-x_center)+(y2-y_center)*(y2-y_center))
	return minValue <= radius*radius
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
