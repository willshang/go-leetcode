package main

import "fmt"

func main() {
	fmt.Println(equalSubstring("abcd", "bcdf", 3))
}

// leetcode1208_尽可能使字符串相等
func equalSubstring(s string, t string, maxCost int) int {
	arr := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = getValue(s[i], t[i])
	}
	left, right := 0, 0
	total := 0
	res := 0
	for right < len(s) {
		for total+arr[right] > maxCost {
			total = total - arr[left]
			left++
		}
		total = total + arr[right]
		if right-left+1 > res {
			res = right - left + 1
		}
		right++
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
