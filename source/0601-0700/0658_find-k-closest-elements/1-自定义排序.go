package main

import "sort"

func main() {

}

func findClosestElements(arr []int, k int, x int) []int {
	sort.Slice(arr, func(i, j int) bool { // 按差值的绝对值排序，相同按照值大小排序
		if abs(arr[i]-x) == abs(arr[j]-x) {
			return arr[i] < arr[j]
		}
		return abs(arr[i]-x) < abs(arr[j]-x)
	})
	res := arr[:k]
	sort.Ints(res)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
