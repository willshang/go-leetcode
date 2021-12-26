package main

import "sort"

func main() {

}

// 返回第一个大于target的位置
func upper_bound(arr []int, target int) int {
	index := sort.SearchInts(arr, target)
	if index == len(arr) || arr[index] != target {
		return index
	}
	for index < len(arr) && arr[index] == target {
		index++
	}
	return index
}

func upper_bound2(arr []int, target int) int {
	return sort.SearchInts(arr, target+1)
}
