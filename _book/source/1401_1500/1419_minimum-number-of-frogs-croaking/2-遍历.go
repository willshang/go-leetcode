package main

import "fmt"

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
	fmt.Println(minNumberOfFrogs("crcoakroak"))
}

func minNumberOfFrogs(croakOfFrogs string) int {
	res := 0
	var c, r, o, a, k int
	for _, char := range croakOfFrogs {
		if char == 'c' {
			c++
		} else if char == 'r' {
			r++
		} else if char == 'o' {
			o++
		} else if char == 'a' {
			a++
		} else if char == 'k' {
			k++
		} else {
			return -1
		}
		res = max(res, c)
		if c < r || r < o || o < a || a < k {
			return -1
		}
		if k == 1 {
			c--
			r--
			o--
			a--
			k--
		}
	}
	if c == 0 && r == 0 && o == 0 && a == 0 && k == 0 {
		return res
	}
	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
