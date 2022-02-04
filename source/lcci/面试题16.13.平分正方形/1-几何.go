package main

func main() {

}

// 程序员面试金典16.13_平分正方形
func cutSquares(square1 []int, square2 []int) []float64 {
	// 2个正方形的中点坐标
	x1, y1, z1 := float64(square1[0])+float64(square1[2])/2, float64(square1[1])+float64(square1[2])/2, float64(square1[2])
	x2, y2, z2 := float64(square2[0])+float64(square2[2])/2, float64(square2[1])+float64(square2[2])/2, float64(square2[2])
	var a, b, c, d float64
	if x1 == x2 { // 1、垂直
		a, b, c, d = x1, min(float64(square1[1]), float64(square2[1])), x1, max(float64(square1[1])+z1, float64(square2[1])+z2)
		return []float64{a, b, c, d}
	}
	// 2、有斜率: y = kx + b1
	k := (y1 - y2) / (x1 - x2)
	b1 := y1 - k*x1
	if abs(k) > 1 { // 斜率大于1，交点通过正方形的上边+下边（根据纵坐标求横坐标）
		b = min(float64(square1[1]), float64(square2[1]))
		d = max(float64(square1[1])+z1, float64(square2[1])+z2)
		a = (b - b1) / k
		c = (d - b1) / k
	} else { // 斜率小于等于1，交点通过正方形的左边+右边（根据横坐标求纵坐标）
		a = min(float64(square1[0]), float64(square2[0]))
		c = max(float64(square1[0])+z1, float64(square2[0])+z2)
		b = a*k + b1
		d = c*k + b1
	}
	if a > c {
		a, c = c, a
		b, d = d, b
	}
	return []float64{a, b, c, d}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a > b {
		return b
	}
	return a
}
