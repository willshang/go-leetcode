package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d ", &n)
	var res = 0
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
			res++
		} else {
			n = (3*n + 1) / 2
			res++
		}
	}
	fmt.Println(res)
}
