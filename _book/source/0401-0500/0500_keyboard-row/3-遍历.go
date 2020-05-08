package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(str))
}

func findWords(words []string) []string {
	res := make([]string, 0, len(words))
	for _, word := range words {
		w := strings.ToLower(word)
		flag := 0
		for _, m := range w {
			switch m {
			case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
				if flag != 0 && flag != 1 {
					flag = 4
					break
				}
				flag = 1
			case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
				if flag != 0 && flag != 2 {
					flag = 4
					break
				}
				flag = 2
			case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
				if flag != 0 && flag != 3 {
					flag = 4
					break
				}
				flag = 3
			default:
				flag = 4
			}
		}
		if flag != 0 && flag != 4 {
			res = append(res, word)
		}
	}
	return res
}
