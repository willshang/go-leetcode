package main

import "strings"

func main() {

}

// leetcode1694_重新格式化电话号码
func reformatNumber(number string) string {
	str := strings.ReplaceAll(number, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	res := ""
	for len(str) > 4 {
		res = res + str[:3] + "-"
		str = str[3:]
	}
	if len(str) == 4 {
		res = res + str[:2] + "-" + str[2:]
	} else {
		res = res + str
	}
	return res
}
