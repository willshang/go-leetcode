package main

import "fmt"

func main() {
	fmt.Println(countBits(10))
}

func countBits(num int) []int {
	res := make([]int, 0)
	for i := 0; i <= num; i++ {
		count := 0
		value := i
		for value != 0 {
			if value%2 == 1 {
				count++
			}
			value = value / 2
		}
		res = append(res, count)
	}
	return res
}
