package main

import (
	"sort"
	"strings"
)

func main() {

}

// 剑指OfferII063.替换单词
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
