package main

import "fmt"

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
}

func equalSubstring(s string, t string, maxCost int) int {
	left := 0
	total := 0
	for right := 0; right < len(s); right++ {
		total = total + getValue(s[right], t[right])
		if total > maxCost {
			total = total - getValue(s[left], t[left])
			left++
		}
	}
	return len(s) - left
}

func getValue(a, b byte) int {
	res := int(a) - int(b)
	if res < 0 {
		return -res
	}
	return res
}
