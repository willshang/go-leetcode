package main

import "fmt"

func main() {
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
}

// leetcode1170_比较字符串最小字母出现频次
func numSmallerByFrequency(queries []string, words []string) []int {
	wordsArr := make([]int, len(words))
	res := make([]int, 0)
	for i := 0; i < len(words); i++ {
		wordsArr[i] = f(words[i])
	}
	for i := 0; i < len(queries); i++ {
		value := f(queries[i])
		count := 0
		for j := 0; j < len(wordsArr); j++ {
			if value < wordsArr[j] {
				count++
			}
		}
		res = append(res, count)
	}
	return res
}

func f(str string) int {
	min := str[0]
	count := 1
	for i := 1; i < len(str); i++ {
		if str[i] < min {
			min = str[i]
			count = 1
		} else if str[i] == min {
			count++
		}
	}
	return count
}
