package main

import (
	"fmt"
)

func main() {
	var str string
	for {
		n, _ := fmt.Scanf("%s", &str)
		if n == 0 {
			break
		}
		for i := len(str) - 1; i >= 0; i-- {
			fmt.Print(string(str[i]))
		}
		fmt.Println()
	}
}
