package main

import "fmt"

func main() {
	target := 9
	for i := 1; i <= target; i++ {
		if target|i == target {
			fmt.Printf("%05b\n", i)
		}
	}
}

/*
00001
01000
01001
*/
