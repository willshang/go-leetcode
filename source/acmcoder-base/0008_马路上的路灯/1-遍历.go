package main

import "fmt"

func main() {
	var M, N int
	fmt.Scan(&M, &N)
	arr := make([][2]int, N)
	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		arr[i] = [2]int{a, b}
	}
	res := getResult(arr, M)
	fmt.Println(res)
}

func getResult(arr [][2]int, M int) int {
	temp := make([]int, M+1)
	for i := 0; i < len(arr); i++ {
		a, b := arr[i][0], arr[i][1]
		for j := a; j <= b; j++ {
			temp[j] = 1
		}
	}
	res := 0
	for i := 0; i <= M; i++ {
		if temp[i] == 0 {
			res++
		}
	}
	return res
}
