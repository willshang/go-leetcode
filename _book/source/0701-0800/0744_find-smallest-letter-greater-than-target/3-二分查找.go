package main

import (
	"fmt"
)

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')

	fmt.Println(string(nextGreatestLetter(letters, target)))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1
	for left <= right {
		mid := left + (right-left)/2
		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return letters[left%len(letters)]
}
