package main

import (
	"fmt"
	"sort"
)

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('a')

	fmt.Println(string(nextGreatestLetter(letters, target)))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)
	i := sort.Search(n, func(i int) bool {
		return target < letters[i]
	})
	return letters[i%n]
}
