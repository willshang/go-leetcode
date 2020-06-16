package main

import (
	"fmt"
	"strconv"
)

func main() {
	//fmt.Println(minNumber([]int{10, 2}))
	//fmt.Println(minNumber([]int{5,3, 30, 34,  9}))
	//fmt.Println(minNumber([]int{1,1,1}))
	fmt.Println(minNumber([]int{1, 4, 7, 2, 5, 8, 0, 3, 6, 9}))

}

var arr []string

func minNumber(nums []int) string {
	arr = make([]string, 0)
	for i := 0; i < len(nums); i++ {
		arr = append(arr, strconv.Itoa(nums[i]))
	}
	quickSort(0, len(arr)-1)
	str := ""
	for i := 0; i < len(arr); i++ {
		str = str + arr[i]
	}
	return str
}

func quickSort(start, end int) {
	if start >= end {
		return
	}
	temp := arr[start]
	i := start
	j := end
	for i < j {
		for i < j && arr[j]+temp >= temp+arr[j] {
			j--
		}
		for i < j && arr[i]+temp <= temp+arr[i] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[start], arr[i] = arr[i], temp
	quickSort(start, i-1)
	quickSort(i+1, end)
}
