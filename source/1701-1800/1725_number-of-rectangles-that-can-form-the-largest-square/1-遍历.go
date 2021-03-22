package main

func main() {

}

// leetcode1725_可以形成最大正方形的矩形数目
func countGoodRectangles(rectangles [][]int) int {
	res := 0
	maxValue := 0
	for i := 0; i < len(rectangles); i++ {
		minValue := min(rectangles[i][0], rectangles[i][1])
		if minValue > maxValue {
			res = 1
			maxValue = minValue
		} else if minValue == maxValue {
			res++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
