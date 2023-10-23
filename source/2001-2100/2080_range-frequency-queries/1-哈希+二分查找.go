package main

import "fmt"

func main() {
	arr := []int{12, 33, 4, 56, 22, 2, 34, 33, 22, 12, 34, 56}
	c := Constructor(arr)
	//fmt.Println(c.Query(1,2,4))
	fmt.Println(c.Query(0, 11, 33))

}

type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]] = append(m[arr[i]], i)
	}
	return RangeFreqQuery{m: m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	arr := this.m[value]
	l := lowerBound(arr, left)
	r := upperBound(arr, right)
	return r - l
}

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
