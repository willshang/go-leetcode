package main

import (
	"fmt"
)

func main() {
	arr := [][]int{
		{0, 1},
		{1, 0},
	}
	fmt.Println(kClosest(arr, 1))
}

func kClosest(points [][]int, K int) [][]int {
	quick(points, 0, len(points)-1, K)
	return points[:K]
}

func quick(points [][]int, left, right int, K int) {
	if left >= right {
		return
	}
	for {
		target := partition(points, left, right)
		if target == K-1 {
			return
		}
		if target < K-1 {
			left = target + 1
		} else {
			right = target - 1
		}
	}
}

func partition(points [][]int, left, right int) int {
	baseValue := points[left]
	for left < right {
		for left < right && dist(baseValue) <= dist(points[right]) {
			right--
		}
		points[left] = points[right]
		for left < right && dist(points[left]) <= dist(baseValue) {
			left++
		}
		points[right] = points[left]
	}
	points[right] = baseValue
	return right
}

func dist(points []int) int {
	return points[0]*points[0] + points[1]*points[1]
}
