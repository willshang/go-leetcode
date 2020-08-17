package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {
	res := 0
	for i := 0; i < len(words); i++ {
		flag := true
		for _, v := range words[i] {
			if strings.Count(words[i], string(v)) > strings.Count(chars, string(v)) {
				flag = false
				continue
			}
		}
		if flag == true {
			res = res + len(words[i])
		}
	}
	return res
}
