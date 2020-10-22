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
		for i := 0; i < len(num); i++ {
			m[num[i]] = 1
		}
		fmt.Println(len(m))
	}
}
