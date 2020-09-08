package main

import "fmt"

func main() {
	fmt.Println(removeDuplicateLetters("bca"))
}

func removeDuplicateLetters(s string) string {
	stack := make([]byte, 0)
	arr := [256]byte{}
	m := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		arr[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == true {
			arr[s[i]]--
			continue
		}
		// arr[栈顶]说明有重复元素
		// 栈顶>s[i]:说明字典序不满足
		for len(stack) > 0 && stack[len(stack)-1] > s[i] && arr[stack[len(stack)-1]] > 0 {
			m[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, s[i])
		arr[s[i]]--
		m[s[i]] = true
	}
	return string(stack)
}
