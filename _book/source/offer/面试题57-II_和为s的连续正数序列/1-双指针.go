package main

import "fmt"

func main() {
	fmt.Println(findContinuousSequence(9))
}

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	i := 1
	j := 2
	for i < j {
		sum := (i + j) * (j - i + 1) / 2
		if sum == target {
			arr := make([]int, 0)
			for k := i; k <= j; k++ {
				arr = append(arr, k)
			}
			res = append(res, arr)
			i++
		} else if sum < target {
			j++
		} else {
			i++
		}
	}
	return res
}
