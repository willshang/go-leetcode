package main

import "strings"

func main() {

}

// leetcode2315_统计星号
func countAsterisks(s string) int {
	res := 0
	arr := strings.Split(s, "|")
	for i := 0; i < len(arr); i = i + 2 {
		res = res + strings.Count(arr[i], "*")
	}
	return res
}
