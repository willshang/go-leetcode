package main

import (
	"fmt"
)

func main() {
	var num string
	for {
		n, _ := fmt.Scanf("%s", &num)
		if n == 0 {
			break
		}
		for i := len(num) - 1; i >= 0; i-- {
			// fmt.Print(num[i] - '0')
			fmt.Print(string(num[i]))
		}
		fmt.Println()
	}
}
