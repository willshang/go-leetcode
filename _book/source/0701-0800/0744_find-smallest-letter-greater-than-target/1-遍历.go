package main

import (
	"fmt"
)

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')

	fmt.Println(string(nextGreatestLetter(letters, target)))
}

// leetcode744_寻找比目标字母大的最小字母
func nextGreatestLetter(letters []byte, target byte) byte {
	for i := 0; i < len(letters); i++ {
		if letters[i] > target {
			return letters[i]
		}
	}
	return letters[0]
}
