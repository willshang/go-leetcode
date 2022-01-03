package main

import "fmt"

func main() {
	target := 7
	for i := target; i > 0; i = (i - 1) & target {
		fmt.Printf("%05b\n", i)
	}
}

/*
00111
00110
00101
00100
00011
00010
00001
*/
