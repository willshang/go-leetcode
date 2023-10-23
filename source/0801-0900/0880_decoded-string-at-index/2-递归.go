package main

import "fmt"

func main() {
	fmt.Println(decodeAtIndex("ha22", 8))
	fmt.Println(decodeAtIndex("leet2code3", 5))
}

func decodeAtIndex(S string, K int) string {
	count := 0 // 字符个数
	n := len(S)
	for i := 0; i < n; i++ { // 先统计最终字符串总长数
		if '0' <= S[i] && S[i] <= '9' {
			value := int(S[i] - '0')
			prev := count
			count = count * value // 重写d-1次，长度是原来的d倍
			if count >= K {
				return decodeAtIndex(S[:i], (K-1)%prev+1) // 缩小范围：避免求余出现0的情况
			}
		} else {
			count++ // 非数字，长度+1
			if count == K {
				return string(S[i])
			}
		}
	}
	return ""
}
