package main

import "sort"

func main() {

}

type node struct {
	char rune
	num  int
}

// leetcode767_重构字符串
func reorganizeString(S string) string {
	arr := make([]node, 26)
	maxCount := 0
	for _, char := range S {
		index := char - 'a'
		arr[index].char = char
		arr[index].num++
		if arr[index].num > maxCount {
			maxCount = arr[index].num
		}
	}
	if maxCount > (len(S)+1)/2 {
		return ""
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].num >= arr[j].num
	})
	res := make([]rune, len(S))
	var index int
	// 先偶后奇
	for i := 0; i < 2; i++ {
		for j := i; j < len(S); j = j + 2 {
			if arr[index].num == 0 {
				index++
			}
			res[j] = arr[index].char
			arr[index].num--
		}
	}
	return string(res)
}
