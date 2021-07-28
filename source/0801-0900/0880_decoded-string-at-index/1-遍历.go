package main

import "fmt"

func main() {
	fmt.Println(decodeAtIndex("ha22", 5))
	fmt.Println(decodeAtIndex("leet2code3", 5))
}

// leetcode880_索引处的解码字符串
func decodeAtIndex(S string, K int) string {
	count := 0 // 字符个数
	n := len(S)
	for i := 0; i < n; i++ { // 先统计最终字符串总长数
		if '0' <= S[i] && S[i] <= '9' {
			value := int(S[i] - '0')
			count = count * value // 重写d-1次，长度是原来的d倍
		} else {
			count++ // 非数字，长度+1
		}
	}
	for i := n - 1; i >= 0; i-- {
		K = K % count                             // 缩小范围
		if K == 0 && 'a' <= S[i] && S[i] <= 'z' { // 找到目标字符
			return string(S[i])
		}
		if '0' <= S[i] && S[i] <= '9' { // 范围缩小
			value := int(S[i] - '0')
			count = count / value
		} else {
			count-- // 长度-1
		}
	}
	return ""
}
