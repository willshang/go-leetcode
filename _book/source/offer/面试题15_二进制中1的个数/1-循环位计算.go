package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(15))
}

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
