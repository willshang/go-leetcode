package main

import "fmt"

func main() {
	fmt.Println(oneEditAway("pale", "ple"))
}

// 程序员面试金典01.05_一次编辑
func oneEditAway(first string, second string) bool {
	if len(first)-len(second) > 1 || len(second)-len(first) > 1 {
		return false
	}
	if first == second {
		return true
	}
	if len(first) > len(second) {
		first, second = second, first
	}
	for i := 0; i < len(first); i++ {
		if first[i] == second[i] {
			continue
		}
		return first[i:] == second[i+1:] || first[i+1:] == second[i+1:]
	}
	return true
}
