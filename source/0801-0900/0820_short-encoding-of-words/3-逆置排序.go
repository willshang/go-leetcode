package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
	fmt.Println(minimumLengthEncoding([]string{"atime", "btime", "aatime"}))
}

func minimumLengthEncoding(words []string) int {
	res := 0
	arr := make([]string, 0)
	for k := range words {
		arr = append(arr, reverse(words[k]))
	}
	sort.Strings(arr)
	for i := 0; i < len(arr)-1; i++ {
		length := len(arr[i])
		if length <= len(arr[i+1]) && arr[i] == arr[i+1][:length] {
			continue
		}
		res = res + length + 1
	}
	return res + len(arr[len(arr)-1]) + 1
}

func reverse(str string) string {
	res := make([]byte, 0)
	for i := len(str) - 1; i >= 0; i-- {
		res = append(res, str[i])
	}
	return string(res)
}
