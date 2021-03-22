package main

import "fmt"

func main() {
	fmt.Println(concatenatedBinary(3))
	fmt.Println(3 & 2)
	fmt.Println(4 & 3)
}

func concatenatedBinary(n int) int {
	res := 0
	length := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			length++
		}
		res = (res<<length + i) % 1000000007
	}
	return res
}
