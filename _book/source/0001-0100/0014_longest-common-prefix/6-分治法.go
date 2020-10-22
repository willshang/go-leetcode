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

	return commonPrefix(strs, 0, len(strs)-1)
}

func commonPrefix(strs []string, left, right int) string {
	if left == right {
		return strs[left]
	}

	middle := (left + right) / 2
	leftStr := commonPrefix(strs, left, middle)
	rightStr := commonPrefix(strs, middle+1, right)
	return commonPrefixWord(leftStr, rightStr)
}

func commonPrefixWord(leftStr, rightStr string) string {
	if len(leftStr) > len(rightStr) {
		leftStr = leftStr[:len(rightStr)]
	}

	if len(leftStr) < 1 {
		return leftStr
	}

	for i := 0; i < len(leftStr); i++ {
		if leftStr[i] != rightStr[i] {
			return leftStr[:i]
		}
	}
	return leftStr
}
