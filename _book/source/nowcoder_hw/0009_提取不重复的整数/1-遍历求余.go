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
		m := make(map[int]int)
		res := 0
		for num > 0 {
			value := num % 10
			num = num / 10
			if m[value] == 0 {
				res = res*10 + value
				m[value] = 1
			}
		}
		fmt.Println(res)
	}
}
