package main

import "math"

func main() {

}

// leetcode1131_绝对值表达式的最大值
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	aMaxValue, bMaxValue, cMaxValue, dMaxValue := math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32
	aMinValue, bMinValue, cMinValue, dMinValue := math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32
	for i := 0; i < len(arr1); i++ {
		aMaxValue = max(aMaxValue, arr1[i]+arr2[i]+i)
		aMinValue = min(aMinValue, arr1[i]+arr2[i]+i)
		bMaxValue = max(bMaxValue, arr1[i]+arr2[i]-i)
		bMinValue = min(bMinValue, arr1[i]+arr2[i]-i)
		cMaxValue = max(cMaxValue, arr1[i]-arr2[i]+i)
		cMinValue = min(cMinValue, arr1[i]-arr2[i]+i)
		dMaxValue = max(dMaxValue, arr1[i]-arr2[i]-i)
		dMinValue = min(dMinValue, arr1[i]-arr2[i]-i)
	}
	a, b := aMaxValue-aMinValue, bMaxValue-bMinValue
	c, d := cMaxValue-cMinValue, dMaxValue-dMinValue
	return max(a, max(b, max(c, d)))
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
