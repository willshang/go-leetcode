package main

import "fmt"

func main() {
	fmt.Println(expressiveWords("heeellooo", []string{"hello", "hi", "helo"}))
}

// leetcode809_情感丰富的文字
func expressiveWords(S string, words []string) int {
	res := 0
	arr := getCount(S)
	for i := 0; i < len(words); i++ {
		temp := getCount(words[i])
		if len(temp) != len(arr) {
			continue
		}
		flag := true
		for j := 0; j < len(arr); j = j + 2 {
			if arr[j] != temp[j] {
				flag = false
				break
			}
			if arr[j+1] == temp[j+1] {
				continue
			}
			if arr[j+1] < 3 || arr[j+1] < temp[j+1] {
				flag = false
				break
			}
		}
		if flag == true {
			res++
		}
	}
	return res
}

func getCount(str string) []int {
	res := make([]int, 0)
	count := 1
	for i := 0; i < len(str); i++ {
		if i == len(str)-1 || str[i] != str[i+1] {
			res = append(res, int(str[i]), count)
			count = 1
		} else {
			count++
		}
	}
	return res
}
