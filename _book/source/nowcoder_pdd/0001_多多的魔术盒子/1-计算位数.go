package main

import "fmt"

func main() {
	var num int
	var n int
	fmt.Scanf("%d", &num)
	for i := 0; i < num; i++ {
		fmt.Scanf("%d", &n)
		res := 1
		for n > 1 {
			res++
			n = n / 2
		}
		fmt.Println(res)
	}
}
