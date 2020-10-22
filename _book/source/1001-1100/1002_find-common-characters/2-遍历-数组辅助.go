package main

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"}))
}

// leetcode1002_查找常用字符
func commonChars(A []string) []string {
	arr := make([][26]int, len(A))
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			arr[i][A[i][j]-'a']++
		}
	}
	res := make([]string, 0)
	for j := 0; j < 26; j++ {
		minValue := arr[0][j]
		for i := 1; i < len(arr); i++ {
			minValue = min(minValue, arr[i][j])
		}
		for minValue > 0 {
			res = append(res, string(j+'a'))
			minValue--
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
