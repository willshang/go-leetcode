package main

import "strings"

func main() {

}

// leetcode2264_字符串中最大的3位相同数字
func largestGoodInteger(num string) string {
	for i := '9'; i >= '0'; i-- {
		target := strings.Repeat(string(i), 3)
		if strings.Contains(num, target) == true {
			return target
		}
	}
	return ""
}
