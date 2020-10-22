package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, 3)
	i := 0
	for {
		if n == 0 {
			break
		}
		arr[i] = n % 10
		n = n / 10
		i++
	}
	for k := 0; k < arr[2]; k++ {
		fmt.Print("B")
	}
	for k := 0; k < arr[1]; k++ {
		fmt.Print("S")
	}
	for k := 0; k < arr[0]; k++ {
		fmt.Print(k + 1)
	}
}
