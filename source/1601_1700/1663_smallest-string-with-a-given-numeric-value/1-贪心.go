package main

import "strings"

func main() {

}

func getSmallestString(n int, k int) string {
	res := ""
	k = k - n
	a := k / 25
	b := k % 25
	right := a
	var left, middle int
	if b == 0 {
		left = n - right
		middle = 0
	} else {
		left = n - right - 1
		middle = b
	}
	res = res + strings.Repeat("a", left)
	if middle > 0 {
		res = res + string('a'+middle)
	}
	res = res + strings.Repeat("z", right)
	return res
}
