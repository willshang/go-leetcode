package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	M = M % N
	arr := make([]int, 0)
	for i := 0; i < N; i++ {
		var num int
		fmt.Scanf("%d", &num)
		arr = append(arr, num)
	}
	if M != 0 {
		reverse(arr)
		reverse(arr[:M])
		reverse(arr[M:])
	}
	for i := 0; i < len(arr)-1; i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
	fmt.Print(arr[len(arr)-1])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
