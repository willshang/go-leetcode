package main

import "sort"

func main() {

}

// leetcode1465_切割后面积最大的蛋糕
func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	hArr := append(horizontalCuts, h, 0)
	wArr := append(verticalCuts, w, 0)
	sort.Ints(hArr)
	sort.Ints(wArr)
	maxHeight := 0
	maxWeight := 0
	for i := 1; i < len(hArr); i++ {
		if hArr[i]-hArr[i-1] > maxHeight {
			maxHeight = hArr[i] - hArr[i-1]
		}
	}
	for i := 1; i < len(wArr); i++ {
		if wArr[i]-wArr[i-1] > maxWeight {
			maxWeight = wArr[i] - wArr[i-1]
		}
	}
	return maxHeight * maxWeight % 1000000007
}
