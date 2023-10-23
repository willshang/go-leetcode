package main

import "fmt"

func main() {
	fmt.Println(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	if n < 10 {
		return n
	}
	arr := make([]int, 50)
	max := 0
	for i := 1; i <= n; i++ {
		value := sum(i)
		arr[value]++
		if arr[value] > max {
			max = arr[value]
		}
	}
	res := 0
	for i := range arr {
		if arr[i] == max {
			res++
		}
	}
	return res
}

func sum(n int) int {
	res := 0
	for n > 0 {
		res = res + n%10
		n = n / 10
	}
	return res
}
