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
		m := make(map[byte]int)
		for i := len(num) - 1; i >= 0; i-- {
			if m[num[i]] == 0 {
				fmt.Printf("%c", num[i])
				m[num[i]] = 1
			}
		}
		fmt.Println()
	}
}
