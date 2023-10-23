package main

import (
	"sort"
	"strings"
)

func main() {

}

// leetcode1859_将句子排序
func sortSentence(s string) string {
	arr := strings.Split(s, " ")
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][len(arr[i])-1] < arr[j][len(arr[j])-1]
	})
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i][:len(arr[i])-1]
	}
	return strings.Join(arr, " ")
}
