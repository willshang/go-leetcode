package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		arr[i] = temp
	}
	res := getResult(arr, m)
	fmt.Println(res)
}

func getResult(arr []int, m int) int {
	res := 1
	count := 0
	for i := 0; i < len(arr); i++ {
		if count+arr[i] <= m {
			count = count + arr[i]
			continue
		} else {
			count = arr[i]
			res++
		}
	}
	return res
}
