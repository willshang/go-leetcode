package main

import "fmt"

func main() {
	fmt.Println(countCharacters([]string{"cat", "bt", "hat", "tree"}, "atach"))
}

func countCharacters(words []string, chars string) int {
	m := make([]int, 26)
	for i := range chars {
		m[chars[i]-'a']++
	}
	res := 0
	for i := 0; i < len(words); i++ {
		temp := make([]int, 26)
		flag := true
		for j := range words[i] {
			temp[words[i][j]-'a']++
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
