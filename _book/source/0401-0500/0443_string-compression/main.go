package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'e', 'e', 'e', 'b', 'b', 'b'}
	fmt.Println(compress(a))
	fmt.Println(compress([]byte{'a'}))
}
func compress(chars []byte) int {
	end, count := 0, 1
	for i, char := range chars {
		if i+1 < len(chars) && char == chars[i+1] {
			count++
		} else {
			chars[end] = char
			end++

			if count > 1 {
				for _, num := range strconv.Itoa(count) {
					chars[end] = byte(num)
					end++
				}
			}
			count = 1
		}
	}
	return end
}
