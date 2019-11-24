package main

import "fmt"

func main() {
	arr := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(arr))
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res = res ^ n
	}
	return res
}
