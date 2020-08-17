package main

import "fmt"

func main() {
	fmt.Println(canPermutePalindrome("tactcoa"))
}

func canPermutePalindrome(s string) bool {
	arr := [256]int{}
	for i := 0; i < len(s); i++ {
		arr[s[i]]++
	}
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			count++
		}
	}
	return count <= 1
}
