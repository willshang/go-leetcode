package main

import (
	"fmt"
)

func main() {
	str := "00110011"
	fmt.Println(countBinarySubstrings(str))
}
func countBinarySubstrings(s string) int {
	count, countZero, countOne := 0,0,0
	prev := rune(s[0])
	for _, r := range s{
		if prev == r{
			if r == '0'{
				countZero++
			}else {
				countOne++
			}
		}else {
			count = count + min(countOne,countZero)
			if r == '0'{
				countZero = 1
			}else {
				countOne = 1
			}
		}
		prev = r
	}
	return count + min(countOne,countZero)
}

func min(a,b int)int  {
	if a < b {
		return a
	}
	return b
}