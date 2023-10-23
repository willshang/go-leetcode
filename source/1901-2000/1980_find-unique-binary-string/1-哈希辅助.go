package main

import (
	"fmt"
	"strings"
)

func main() {

}

// leetcode1980_找出不同的二进制字符串
func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	m := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		m[strings.Repeat("0", 16-n)+nums[i]] = true
	}
	for i := 0; i < (1 << n); i++ {
		if m[fmt.Sprintf("%016b", i)] == false {
			res := fmt.Sprintf("%016b", i)
			return res[16-n:]
		}
	}
	return ""
}
