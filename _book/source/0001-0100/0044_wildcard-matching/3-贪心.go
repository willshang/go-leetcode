package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
}

func isMatch(s string, p string) bool {
	i, j := 0, 0
	start, last := 0, 0
	for i = 0; i < len(s); {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
		} else if j < len(p) && p[j] == '*' {
			last = i      // 记录s下一个的位置
			start = j + 1 // 记录*下一个的位置
			j++
		} else if start != 0 {
			last++
			i = last // 更新到s记录位置的下一个
			j = start
		} else {
			return false
		}
	}
	for ; j < len(p) && p[j] == '*'; j++ {
	}
	return j == len(p)
}
