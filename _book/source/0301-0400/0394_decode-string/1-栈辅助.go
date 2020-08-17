package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

// leetcode394_字符串解码
func decodeString(s string) string {
	res := make([]byte, 0)
	numStack := make([]int, 0)
	lenStack := make([]int, 0)
	var count int
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			count = count*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, count)
			count = 0
			lenStack = append(lenStack, len(res))
		} else if s[i] == ']' {
			c := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			l := lenStack[len(lenStack)-1]
			lenStack = lenStack[:len(lenStack)-1]
			str := res[l:]
			res = res[:l]
			for i := 0; i < c; i++ {
				res = append(res, str...)
			}
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
