package main

import "fmt"

func main() {
	fmt.Println(numberOfSteps(14))
}

func numberOfSteps(num int) int {
	res := 0
	for num > 0 {
		if num%2 == 1 {
			num = num - 1
		} else {
			num = num / 2
		}
		res++
	}
	return res
}
