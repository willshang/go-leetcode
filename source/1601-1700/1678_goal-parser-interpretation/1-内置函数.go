package main

import "strings"

func main() {

}

// leetcode1678_设计Goal解析器
func interpret(command string) string {
	command = strings.ReplaceAll(command, "(al)", "al")
	command = strings.ReplaceAll(command, "()", "o")
	return command
}
