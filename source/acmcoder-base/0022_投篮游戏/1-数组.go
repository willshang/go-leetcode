package main

import "fmt"

func main() {
	var p, n int
	fmt.Scan(&p, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		arr[i] = temp
	}
	res := getResult(arr, p)
	fmt.Println(res)
}

func getResult(arr []int, p int) int {
	m := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		target := arr[i] % p
		if m[target] == true {
			return i + 1
		}
		m[target] = true
	}
	return -1
}
