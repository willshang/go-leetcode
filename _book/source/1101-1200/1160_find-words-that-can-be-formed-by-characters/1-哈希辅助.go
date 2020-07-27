package main

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

// leetcode1160_拼写单词
func countCharacters(words []string, chars string) int {
	m := make(map[byte]int)
	for i := range chars {
		m[chars[i]]++
	}
	res := 0
	for i := 0; i < len(words); i++ {
		temp := make(map[byte]int)
		flag := true
		for j := range words[i] {
			temp[words[i][j]]++
		}
		if len(temp) > len(m) {
			continue
		}
		for k, v := range temp {
			if v > m[k] {
				flag = false
				break
			}
		}
		if flag == true {
			res = res + len(words[i])
		}
	}
	return res
}
