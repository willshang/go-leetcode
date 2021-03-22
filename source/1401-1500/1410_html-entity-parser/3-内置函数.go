package main

import (
	"fmt"
)

func main() {
	//fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))
	//fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	fmt.Println(entityParser("&amp;amp;amp;gt;"))
}

func entityParser(text string) string {
	res := make([]byte, 0)
	temp := make([]byte, 0)
	for i := 0; i < len(text); i++ {
		if text[i] == '&' {
			if len(temp) == 0 {
				temp = append(temp, text[i])
			} else {
				res = append(res, temp...)
				temp = make([]byte, 0)
				temp = append(temp, text[i])
			}
		} else if text[i] == ';' {
			temp = append(temp, text[i])
			switch string(temp) {
			case "&gt;":
				res = append(res, '>')
			case "&lt;":
				res = append(res, '<')
			case "&quot;":
				res = append(res, '"')
			case "&apos;":
				res = append(res, '\'')
			case "&frasl;":
				res = append(res, '/')
			case "&amp;":
				res = append(res, '&')
			default:
				res = append(res, temp...)
			}
			temp = make([]byte, 0)
		} else if len(temp) == 0 {
			res = append(res, text[i])
		} else {
			temp = append(temp, text[i])
		}
	}
	if len(temp) > 0 {
		res = append(res, temp...)
	}
	return string(res)
}
