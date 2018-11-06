package main

import (
	"fmt"
)

func main() {
	fmt.Println(romanToInt("LIII"))
}

func romanToInt(s string) int {
	result := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		current := m[s[i]]
		flag := 1
		if current < last {
			flag = -1
		}
		result = result + flag * current
		last = current
	}
	return result
}
