package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(binaryGap(22))
}

func binaryGap(N int) int {
	res := 0
	str := strconv.FormatInt(int64(N), 2)
	j := -1
	for i := 0; i < len(str); i++ {
		if str[i] == '1' {
			if j == -1 {
				j = i
			} else {
				if i-j > res {
					res = i - j
				}
				j = i
			}
		}
	}
	return res
}
