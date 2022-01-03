package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10}
	fmt.Println(len(arr))
	fmt.Println(leftSearch(arr, 7))
	fmt.Println(rightSearch(arr, 7))
	fmt.Println(upperBound(arr, 9))
	fmt.Println(upperBound(arr, 11))
}

// 1、寻找目标值
// 寻找目标值
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 左边界：找不到返回-1
func leftSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid - 1 // 相等情况，左移
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(arr) || arr[left] != target { // 越界或者找不到
		return -1
	}
	return left
}

// 右边界：找不到返回-1
func rightSearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1 // 相等情况，右移
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || arr[right] != target { // 越界或者找不到
		return -1
	}
	return right
}

// 2、查找边界
// 返回第一个大于等于target的位置
func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid // 收缩左边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// 返回第一个大于target的位置
func upperBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1 // 收缩左边界
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
