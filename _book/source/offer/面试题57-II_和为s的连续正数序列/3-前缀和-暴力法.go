package main

import "fmt"

func main() {
	fmt.Println(findContinuousSequence(9))
}

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	arr := make([]int, target+1)
	for i := 1; i <= target; i++ {
		arr[i] = arr[i-1] + i
	}
	for i := 1; i <= (target+1)/2; i++ {
		for j := i + 1; j <= target && arr[j]-arr[i-1] <= target; j++ {
			if arr[j]-arr[i-1] == target {
				temp := make([]int, 0)
				for k := i; k <= j; k++ {
					temp = append(temp, k)
				}
				res = append(res, temp)
				break
			}
		}
	}
	return res
}
