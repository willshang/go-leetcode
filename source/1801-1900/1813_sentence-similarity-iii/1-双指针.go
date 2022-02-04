package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(areSentencesSimilar("A", "a A a A"))
	//fmt.Println(areSentencesSimilar("My name is Haley", "My Haley"))
	//fmt.Println(areSentencesSimilar("Eating right now", "Eating"))
	fmt.Println(areSentencesSimilar("a b c d e", "b d"))
}

// leetcode1813_句子相似性III
func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if len(sentence1) > len(sentence2) { // 交换保证 1 < 2
		sentence1, sentence2 = sentence2, sentence1
	}
	// 4种情况：
	// 1、全等： AB = AB
	// 2、头部等于：AXX = A
	// 3、尾部等于：XXA = A
	// 4、首尾相等：AXXB = AB
	a := strings.Fields(sentence1)
	b := strings.Fields(sentence2)
	i, j := 0, 0
	for i < len(a) && a[i] == b[i] { // 从开头统计
		i++
	}
	for 0 <= len(a)-1-j && a[len(a)-1-j] == b[len(b)-1-j] { // 从结尾统计
		j++
	}
	return i+j >= len(a) // 大于等于较短的长度
}
