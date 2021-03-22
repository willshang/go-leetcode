package main

import "fmt"

func main() {
	fmt.Println(numOfSubarrays([]int{1, 3, 5}))
}

func numOfSubarrays(arr []int) int {
	odd, even := 0, 1
	res := 0
	total := 0
	for i := 0; i < len(arr); i++ {
		total = total + arr[i]
		if total%2 == 0 {
			res = res + odd
			even = even + 1
		} else {
			res = res + even
			odd = odd + 1
		}
	}
	return res % 1000000007
}
