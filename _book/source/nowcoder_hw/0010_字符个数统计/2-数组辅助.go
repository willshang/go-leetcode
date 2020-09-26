package main

import (
	"fmt"
)

func main() {
	var num string
	for {
		n, _ := fmt.Scan(&num)
		if n == 0 {
			break
		}
		m := make([]int, 128)
		res := 0
		for i := 0; i < len(num); i++ {
			if m[num[i]] == 0 {
				m[num[i]] = 1
				res++
			}
		}
		fmt.Println(res)
	}
}
