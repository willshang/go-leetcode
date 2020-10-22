package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordsArr := make([]int, len(words))
	res := make([]int, 0)
	for i := 0; i < len(words); i++ {
		wordsArr[i] = f(words[i])
	}
	sort.Ints(wordsArr)
	for i := 0; i < len(queries); i++ {
		value := f(queries[i])
		count := binarySearch(value, wordsArr)
		res = append(res, count)
	}
	return res
}

func binarySearch(value int, target []int) int {
	left := 0
	right := len(target) - 1
	for left < right {
		mid := left + (right-left)/2
		if target[mid] > value {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if target[left] <= value {
		return 0
	}
	return len(target) - left
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
