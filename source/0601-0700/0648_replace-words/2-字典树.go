package main

import (
	"strings"
)

func main() {

}

func replaceWords(dictionary []string, sentence string) string {
	trie := Constructor()
	for i := 0; i < len(dictionary); i++ {
		trie.Insert(dictionary[i])
	}
	arr := strings.Split(sentence, " ")
	for i := 0; i < len(arr); i++ {
		result := trie.Search(arr[i])
		if result != "" {
			arr[i] = result
		}
	}
	return strings.Join(arr, " ")
}

type Trie struct {
	next   [26]*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
	ending int       // 次数（可以改为bool）
}

func Constructor() Trie {
	return Trie{
		next:   [26]*Trie{},
		ending: 0,
	}
}

// 插入word
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next:   [26]*Trie{},
				ending: 0,
			}
		}
		temp = temp.next[value]
	}
	temp.ending++
}

// 查找
func (this *Trie) Search(word string) string {
	temp := this
	res := ""
	for _, v := range word {
		res = res + string(v)
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return ""
		}
		if temp.ending > 0 {
			return res
		}
	}
	return ""
}
