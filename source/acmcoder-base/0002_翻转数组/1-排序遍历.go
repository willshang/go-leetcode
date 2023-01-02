package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scanf("%d", &temp)
		arr[i] = temp
	}
	res := getResult(arr)
	if res == true {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

func getResult(arr []int) bool {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(temp)
	left, right := 0, len(arr)-1
	for left < len(arr) && arr[left] == temp[left] {
		left++
	}
	if left == len(arr) {
		return false
	}
	for 0 <= right && arr[right] == temp[right] {
		right--
	}
	for i := left; i < right; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}
	return true
}
