package main

import (
	"fmt"
)

func main() {
	var num int
	for {
		n, _ := fmt.Scanf("%d", &num)
		if n == 0 {
			break
		}
		for num > 0{
			fmt.Print(num%10)
			num = num/10
		}
		fmt.Println()
	}
}
