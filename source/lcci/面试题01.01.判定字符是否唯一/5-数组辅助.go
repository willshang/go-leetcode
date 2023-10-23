package main

import "fmt"

func main() {
	fmt.Println(isUnique("leetcode"))
}

func isUnique(astr string) bool {
	arr := make([]int, 256)
	for i := 0; i < len(astr); i++ {
		if arr[astr[i]] > 0 {
			return false
		}
		arr[astr[i]] = 1
	}
	return true
}
