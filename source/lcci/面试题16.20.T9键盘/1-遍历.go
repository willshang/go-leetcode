package main

import "fmt"

func main() {
	fmt.Println(getValidT9Words("8733", []string{"tree", "used"}))
}

// 程序员面试金典16.20_T9键盘
var m [26]byte = [26]byte{
	'2', '2', '2',
	'3', '3', '3',
	'4', '4', '4',
	'5', '5', '5',
	'6', '6', '6',
	'7', '7', '7', '7',
	'8', '8', '8',
	'9', '9', '9', '9',
}

func getValidT9Words(num string, words []string) []string {
	res := make([]string, 0)
	for _, str := range words {
		if len(str) != len(num) {
			continue
		}
		flag := true
		for i := 0; i < len(str); i++ {
			if num[i] != m[str[i]-'a'] {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, str)
		}
	}
	return res
}
