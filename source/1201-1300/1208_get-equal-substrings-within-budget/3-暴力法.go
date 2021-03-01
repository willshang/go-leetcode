package main

import "fmt"

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
}

func equalSubstring(s string, t string, maxCost int) int {
	res := 0
	for i := 0; i < len(s); i++ {
		total := 0
		count := 0
		for j := i; j < len(s); j++ {
			total = total + getValue(s[j], t[j])
			if total > maxCost {
				break
			}
			count++
		}
		if count > res {
			res = count
		}
	}
	return res
}

func getValue(a, b byte) int {
	res := int(a) - int(b)
	if res < 0 {
		return -res
	}
	return res
}
