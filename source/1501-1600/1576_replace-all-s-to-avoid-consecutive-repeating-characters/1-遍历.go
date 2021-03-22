package main

import "fmt"

func main() {
	fmt.Println(modifyString("?zs"))
}

// leetcode1576_替换所有的问号
func modifyString(s string) string {
	res := []byte(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			for j := byte('a'); j <= 'z'; j++ {
				if (i == 0 || res[i-1] != j) && (i == len(s)-1 || s[i+1] != j) {
					res[i] = j
					break
				}
			}
		}
	}
	return string(res)
}
