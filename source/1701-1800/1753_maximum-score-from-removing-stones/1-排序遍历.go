package main

import "sort"

func main() {

}

func maximumScore(a int, b int, c int) int {
	res := 0
	arr := []int{a, b, c}
	for {
		sort.Ints(arr)
		if arr[2] > 0 && arr[1] > 0 {
			arr[2]--
			arr[1]--
			res++
		} else {
			break
		}
	}
	return res
}
