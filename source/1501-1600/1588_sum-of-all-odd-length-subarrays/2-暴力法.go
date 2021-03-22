package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
	fmt.Println(sumOddLengthSubarrays([]int{1, 2}))
	fmt.Println(sumOddLengthSubarrays([]int{10, 11, 12}))
}

func sumOddLengthSubarrays(arr []int) int {
	res := 0
	for i := 1; i <= len(arr); i = i + 2 {
		for j := 0; j < len(arr); j++ {
			k := j + i
			for start := j; start < k && k <= len(arr); start++ {
				res = res + arr[start]
			}
		}
	}
	return res
}
