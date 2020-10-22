package main

import (
	"fmt"
)

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	arr := [256]int{}
	for i := 0; i < len(s1); i++ {
		arr[s1[i]]++
		arr[s2[i]]--
	}
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}
