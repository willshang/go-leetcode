package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	arr := []byte(s)
	i := len(arr) - 1
	j := len(arr)
	res := ""
	flag := false
	for i >= 0 {
		if arr[i] == ' ' {
			if flag == true {
				res = res + " " + string(arr[i+1:j])
				flag = false
			}
			j = i
			i--
		} else {
			if i == 0 {
				res = res + " " + string(arr[i:j])
			}
			i--
			flag = true
		}
	}
	if len(res) > 0 {
		return res[1:]
	}
	return res
}
