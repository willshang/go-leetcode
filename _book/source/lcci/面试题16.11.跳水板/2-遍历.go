package main

import "fmt"

func main() {
	fmt.Println(divingBoard(1, 2, 3))
}

func divingBoard(shorter int, longer int, k int) []int {
	res := make([]int, 0)
	if k == 0 {
		return res
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	start := shorter * k
	diff := longer - shorter
	for i := 0; i <= k; i++ {
		res = append(res, start+i*diff)
	}
	return res
}
