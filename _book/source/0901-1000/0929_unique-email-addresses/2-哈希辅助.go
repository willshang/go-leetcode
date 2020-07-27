package main

import (
	"fmt"
)

func main() {
	//fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com",
	//	"test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))

	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.email@leetcode.com"}))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)
	for i := 0; i < len(emails); i++ {
		addr := ""
		isBreak := false
		for j := 0; j < len(emails[i]); j++ {
			if emails[i][j] == '+' {
				isBreak = true
			} else if emails[i][j] == '.' {
				continue
			} else if emails[i][j] == '@' {
				addr = addr + emails[i][j:]
				break
			} else if isBreak == true {
			} else {
				addr = addr + string(emails[i][j])
			}
		}
		m[addr] = true
	}
	return len(m)
}
