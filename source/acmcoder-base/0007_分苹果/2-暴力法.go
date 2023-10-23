package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	res := getResult(n)
	fmt.Println(res)
}

func getResult(n int) int {
	for i := 1; ; i++ {
		target := i
		bear := n
		for bear > 0 {
			if (target-1)%n == 0 {
				bear--
				target = target - 1 - (target-1)/n
			} else {
				break
			}
		}
		if bear == 0 {
			return i
		}
	}
	return 0
}
