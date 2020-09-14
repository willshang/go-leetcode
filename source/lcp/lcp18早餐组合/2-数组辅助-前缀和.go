package main

import (
	"fmt"
)

func main() {
	fmt.Println(breakfastNumber([]int{10, 20, 5}, []int{5, 5, 2}, 15))
	fmt.Println(breakfastNumber([]int{2, 1, 1}, []int{8, 9, 5, 1}, 9))
}

func breakfastNumber(staple []int, drinks []int, x int) int {
	res := 0
	arr := make([]int, x+1)
	for i := 0; i < len(staple); i++ {
		if staple[i] < x {
			arr[staple[i]]++
		}
	}
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i]
	}
	for i := 0; i < len(drinks); i++ {
		left := x - drinks[i]
		if left <= 0 {
			continue
		}
		res = (res + arr[left]) % 1000000007
	}
	return res
}
