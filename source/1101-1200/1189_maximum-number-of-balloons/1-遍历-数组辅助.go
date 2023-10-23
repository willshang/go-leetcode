package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))
}

func maxNumberOfBalloons(text string) int {
	m := make([]int, 26)
	str := "ablon"
	for i := 0; i < len(str); i++ {
		m[str[i]-'a']++
	}
	for i := 0; i < len(text); i++ {
		if m[text[i]-'a'] > 0 {
			m[text[i]-'a']++
		}
	}
	min := math.MaxInt32
	for k, v := range m {
		if v == 0 {
			continue
		}
		if k+'a' == 'l' || k+'a' == 'o' {
			v = (v - 1) / 2
		} else {
			v = v - 1
		}
		if v < min {
			min = v
		}
	}
	return min
}
