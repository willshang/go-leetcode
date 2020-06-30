package main

import "fmt"

func main() {
	var num int
	for {
		n, _ := fmt.Scanf("%d", &num)
		if n == 0 {
			break
		}
		res := 0
		for num > 0 {
			if num%2 == 1 {
				res++
			}
			num = num / 2
		}
		fmt.Println(res)
	}
}
