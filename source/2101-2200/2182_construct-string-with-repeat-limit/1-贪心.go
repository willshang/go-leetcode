package main

import "bytes"

func main() {

}

// leetcode2182_构造限制重复的字符串
func repeatLimitedString(s string, repeatLimit int) string {
	arr := make([]int, 26)
	for i := 0; i < len(s); i++ {
		arr[int(s[i]-'a')]++
	}
	res := make([]byte, 0)
	for i := 25; i >= 0; i-- {
		prev := i - 1
		for {
			count := min(repeatLimit, arr[i])
			arr[i] = arr[i] - count
			res = append(res, bytes.Repeat([]byte{byte('a' + i)}, count)...)
			if arr[i] == 0 { // 用完退出
				break
			}
			for prev >= 0 && arr[prev] == 0 {
				prev--
			}
			if prev < 0 { // 没有次大值退出
				break
			}
			arr[prev]--
			res = append(res, byte('a'+prev)) // 添加次大值
		}
	}
	return string(res)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
