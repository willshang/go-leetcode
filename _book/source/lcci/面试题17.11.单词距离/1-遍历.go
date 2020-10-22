package main

import "fmt"

func main() {
	arr := []string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"}
	fmt.Println(findClosest(arr, "a", "student"))
}

// 程序员面试金典17.11_单词距离
func findClosest(words []string, word1 string, word2 string) int {
	res := len(words) - 1
	a, b := -1, -1
	for i := 0; i < len(words); i++ {
		if words[i] == word1 {
			a = i
		}
		if words[i] == word2 {
			b = i
		}
		if a != -1 && b != -1 && abs(a, b) < res {
			res = abs(a, b)
		}
	}
	return res
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
