package main

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("bca"))
}

// leetcode1081_不同字符的最小子序列
func smallestSubsequence(text string) string {
	stack := make([]byte, 0)
	arr := [256]byte{}
	m := make(map[byte]bool)
	for i := 0; i < len(text); i++ {
		arr[text[i]]++
	}
	for i := 0; i < len(text); i++ {
		if m[text[i]] == true {
			arr[text[i]]--
			continue
		}
		// arr[栈顶]说明有重复元素
		// 栈顶>s[i]:说明字典序不满足
		for len(stack) > 0 && stack[len(stack)-1] > text[i] && arr[stack[len(stack)-1]] > 0 {
			m[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, text[i])
		arr[text[i]]--
		m[text[i]] = true
	}
	return string(stack)
}
