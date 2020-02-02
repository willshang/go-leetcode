package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	fmt.Println(longestCommonPrefix([]string{"abc", "acb", "bac"}))

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	commonStr := common(strs[0], strs[1])
	if commonStr == "" {
		return ""
	}
	for i := 2; i < len(strs); i++ {
		if commonStr == "" {
			return ""
		}
		commonStr = common(commonStr, strs[i])
	}
	return commonStr
}

func common(str1, str2 string) string {
	length := 0
	for i := 0; i < len(str1); i++ {
		char := str1[i]
		if i >= len(str2) || char != str2[i] {
			return str1[:length]
		}
		length++
	}
	return str1[:length]
}
