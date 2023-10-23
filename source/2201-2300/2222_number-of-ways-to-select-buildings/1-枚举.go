package main

import "strings"

func main() {

}

func numberOfWays(s string) int64 {
	res := int64(0)
	n := len(s)
	count1 := strings.Count(s, "1") // 1的总个数
	count0 := n - count1            // 0的总个数
	count := 0                      // 遍历时候1的个数
	for i := 0; i < n; i++ {        // 枚举以当前字母作为中间元素
		if s[i] == '1' { // 010
			left0 := i - count       // 左边0的个数
			right0 := count0 - left0 // 右边0的个数
			res = res + int64(left0)*int64(right0)
			count++
		} else { // 101
			left1 := count           // 左边1的个数
			right1 := count1 - count // 右边1的个数
			res = res + int64(left1)*int64(right1)
		}
	}
	return res
}
