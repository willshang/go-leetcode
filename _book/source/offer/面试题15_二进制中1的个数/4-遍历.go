package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(15))
}

func hammingWeight(num uint32) int {
	count := 0
	flag := uint32(1)
	for flag != 0 {
		if num&flag == flag {
			count++
		}
		flag = flag << 1
	}
	return count
}
