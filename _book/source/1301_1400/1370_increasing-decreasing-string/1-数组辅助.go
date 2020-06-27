package main

import "fmt"

func main() {
	fmt.Println(sortString("aaaabbbbcccc"))
}

// leetcode1370_上升下降字符串
func sortString(s string) string {
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[s[i]-'a']++
	}
	res := ""
	for len(res) < len(s) {
		for i := 0; i < 26; i++ {
			if arr[i] > 0 {
				res = res + string(i+'a')
				arr[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if arr[i] > 0 {
				res = res + string(i+'a')
				arr[i]--
			}
		}
	}
	return res
}
