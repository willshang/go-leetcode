package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}

// leetcode929_go_独特的电子邮件地址
func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for i := 0; i < len(emails); i++ {
		addr := ""
		arr := strings.Split(emails[i], "@")
		for j := 0; j < len(arr[0]); j++ {
			if arr[0][j] == '+' {
				break
			} else if arr[0][j] == '.' {
				continue
			} else {
				addr = addr + string(arr[0][j])
			}
		}
		m[addr+"@"+arr[1]] = true
	}
	return len(m)
}
