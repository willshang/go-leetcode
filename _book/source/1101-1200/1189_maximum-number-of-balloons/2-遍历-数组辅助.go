package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}

// leetcode1189_气球的最大数量
func maxNumberOfBalloons(text string) int {
	m := make([]int, 26)
	for i := 0; i < len(text); i++ {
		m[text[i]-'a']++
	}
	res := 0
	str := "balloon"
	for {
		for i := 0; i < len(str); i++ {
			m[str[i]-'a']--
			if m[str[i]-'a'] < 0 {
				return res
			}
		}
		res++
	}
}
