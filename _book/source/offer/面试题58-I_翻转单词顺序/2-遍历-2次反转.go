package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	arr := []byte(s)
	arr = reverse(arr)
	i := 0
	j := 0
	res := ""
	flag := false
	for i < len(arr) {
		if arr[i] == ' ' {
			if flag == true {
				res = res + " " + string(reverse(arr[j:i]))
				flag = false
			}
			i++
			j = i
		} else {
			i++
			if i == len(arr) {
				res = res + " " + string(reverse(arr[j:i]))
			}
			flag = true
		}
	}
	if len(res) > 0 {
		return res[1:]
	}
	return res
}

func reverse(arr []byte) []byte {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	return arr
}
