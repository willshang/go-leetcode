package main

import (
	"sort"
	"strings"
)

func main() {

}

// leetcode648_单词替换
func replaceWords(dictionary []string, sentence string) string {
	sort.Strings(dictionary)
	arr := strings.Split(sentence, " ")
	for i := 0; i < len(arr); i++ {
		for _, v := range dictionary {
			if strings.HasPrefix(arr[i], v) {
				arr[i] = v
				break
			}
		}
	}
	return strings.Join(arr, " ")
}
