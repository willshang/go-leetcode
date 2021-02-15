package main

import "strings"

func main() {

}

// leetcode777_在LR字符串中交换相邻字符
func canTransform(start string, end string) bool {
	if strings.ReplaceAll(start, "X", "") != strings.ReplaceAll(end, "X", "") {
		return false
	}
	j := 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'L' { // LX=>XL, L是往右的
			for end[j] != 'L' {
				j++
			}
			if i < j {
				return false
			}
			j++
		}
	}
	j = 0
	for i := 0; i < len(start); i++ {
		if start[i] == 'R' { // XR=>RX, R是往左的
			for end[j] != 'R' {
				j++
			}
			if i > j {
				return false
			}
			j++
		}
	}
	return true
}
