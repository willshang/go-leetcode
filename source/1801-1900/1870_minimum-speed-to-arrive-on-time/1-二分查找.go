package main

import "math"

func main() {

}

// leetcode1870_准时到达的列车最小时速
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if float64(n-1) > hour {
		return -1
	}
	left, right := 1, math.MaxInt32
	for left <= right {
		mid := left + (right-left)/2
		if judge(dist, float64(mid)) <= hour {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(arr []int, speed float64) float64 {
	n := len(arr)
	res := float64(0)
	for i := 0; i < n-1; i++ {
		res = res + math.Ceil(float64(arr[i])/speed) // 向上取整
	}
	res = res + float64(arr[n-1])/speed // 注意：最后一个不需要等待
	return res
}
