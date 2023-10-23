package main

import (
	"fmt"
)

func main() {
	fmt.Println(replaceSpaces("Mr John Smith    ", 13))
}

// 程序员面试金典01.03_URL化
func replaceSpaces(S string, length int) string {
	res := make([]byte, 0)
	for i := 0; i < length; i++ {
		if S[i] == ' ' {
			res = append(res, '%')
			res = append(res, '2')
			res = append(res, '0')
		} else {
			res = append(res, S[i])
		}
	}
	return string(res)
}
