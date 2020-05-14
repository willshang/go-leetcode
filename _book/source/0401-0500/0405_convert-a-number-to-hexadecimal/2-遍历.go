package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(toHex(-254))
}

var h = []string{
	"0", "1", "2", "3", "4", "5", "6", "7",
	"8", "9", "a", "b", "c", "d", "e", "f",
}

func toHex(num int) string {
	res := ""
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num = num + 4294967296
	}

	for num != 0 {
		temp := num % 16
		res = h[temp] + res
		num = num / 16
	}
	return res
}
