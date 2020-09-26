package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'e', 'e', 'e', 'b', 'b', 'b', 'a'}
	fmt.Println(compress(a))
	fmt.Println(compress([]byte{'a'}))
}

// leetcode443_压缩字符串
func compress(chars []byte) int {
	j := 0
	count := 1
	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if i+1 < len(chars) && char == chars[i+1] {
			count++
		} else {
			chars[j] = char
			j++
			if count > 1 {
				for _, num := range strconv.Itoa(count) {
					chars[j] = byte(num)
					j++
				}
			}
			count = 1
		}
	}
	return j
}
