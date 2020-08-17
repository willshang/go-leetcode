package main

import "fmt"

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

func decodeString(s string) string {
	res := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if '0' <= s[i] && s[i] <= '9' {
			count = count*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			times := 0
			i++
			str := make([]byte, 0)
			for s[i] != ']' || times != 0 {
				if s[i] == '[' {
					times++
				} else if s[i] == ']' {
					times--
				}
				str = append(str, s[i])
				i++
			}
			temp := decodeString(string(str))
			for j := 0; j < count; j++ {
				res = res + (temp)
			}
			count = 0
		} else {
			res = res + string(s[i])
		}
	}
	return res
}
