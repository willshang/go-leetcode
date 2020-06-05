package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}

// leetcode551_go_学生出勤记录 I
func checkRecord(s string) bool {
	aNum := 0
	lNum := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			aNum++
		}
		if s[i] == 'L' {
			lNum++
		} else {
			lNum = 0
		}
		if aNum == 2 {
			return false
		}
		if lNum == 3 {
			return false
		}
	}
	return true
}
