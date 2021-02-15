package main

import "fmt"

func main() {
	fmt.Println(prevPermOpt1([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println(prevPermOpt1([]int{6, 5, 4, 1}))
	fmt.Println(prevPermOpt1([]int{6, 5, 4, 1, 2, 3, 3}))
}

func prevPermOpt1(arr []int) []int {
	a := -1
	b := -1
	var i int
	for i = len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] { // 找到第一个降序的位置a
			a = i
			break
		}
	}
	if a == -1 {
		return arr
	}
	maxValue := -1
	for j := i + 1; j < len(arr); j++ { // 往后找到小于的最大数第1次出现位置b
		if arr[a] > arr[j] {
			if arr[j] > maxValue {
				maxValue = arr[j]
				b = j
			}
		}
	}
	if a != -1 && b != -1 {
		arr[a], arr[b] = arr[b], arr[a]
	}
	return arr
}
