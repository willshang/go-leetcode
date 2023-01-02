package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	res := getResult(str)
	fmt.Println(res)
}

func getResult(str string) string {
	for i := 0; i < len(str); i++ {
		if isPalindrome(str[0:i]+str[i+1:]) == true {
			return "Yes" // 注意跟题目大小写不一样
		}
	}
	return "No" // 注意跟题目大小写不一样
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
