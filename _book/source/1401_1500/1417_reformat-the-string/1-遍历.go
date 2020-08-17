package main

import "fmt"

func main() {
	//fmt.Println(reformat("a0b1c2"))
	fmt.Println(reformat("covid2019"))
}

func reformat(s string) string {
	arr := make([]byte, 0)
	str := make([]byte, 0)
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			arr = append(arr, s[i])
		} else {
			str = append(str, s[i])
		}
	}
	if abs(len(arr), len(str)) > 1 {
		return res
	} else {
		length := len(arr)
		if len(str) < length {
			length = len(str)
		}
		for i := 0; i < length; i++ {
			res = res + string(arr[i])
			res = res + string(str[i])
		}
		if length == len(str) && length < len(arr) {
			res = res + string(arr[length])
		} else if length == len(arr) && length < len(str) {
			res = string(str[length]) + res
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
