package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(areSentencesSimilar("A", "a A a A"))
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	var a, b []string
	a = strings.Fields(sentence1)
	b = strings.Fields(sentence2)
	if len(a) > len(b) {
		a, b = b, a
	}
	arr := make([]int, len(b))
	j := 0
	flag := false
	for i := 0; i < len(b); i++ {
		if j < len(a) && a[j] == b[i] {
			arr[i] = 1
			j++
			flag = true
		}
	}
	// 没有1个相等 或者 较长者遍历完而较短没有
	if flag == false || j < len(a) {
		return false
	}
	fmt.Println(arr)
	left := 0
	right := len(arr) - 1
	for left < len(arr) && arr[left] == 1 {
		left++
	}
	for 0 <= right && arr[right] == 1 {
		right--
	}
	// 中间只能存在0的情况
	for i := left; i <= right; i++ {
		if arr[i] == 1 {
			return false
		}
	}
	return true
}
