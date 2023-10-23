package main

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"a"}, "a"))
}

func camelMatch(queries []string, pattern string) []bool {
	node := Constructor()
	node.Insert(pattern)
	res := make([]bool, len(queries))
	for i := 0; i < len(queries); i++ {
		res[i] = node.Match(queries[i])
	}
	return res
}

type Trie struct {
	next   map[int32]*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
	ending int             // 次数（可以改为bool）
}

func Constructor() *Trie {
	return &Trie{
		next:   make(map[int32]*Trie),
		ending: 0,
	}
}

// 插入word
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v
		if _, ok := temp.next[value]; ok == false {
			temp.next[value] = Constructor()
		}
		temp = temp.next[value]
	}
	temp.ending++
}

// 查找
func (this *Trie) Match(word string) bool {
	temp := this
	for _, v := range word {
		value := v
		if _, ok := temp.next[value]; ok == false {
			if value < 'a' {
				return false
			}
		} else {
			temp = temp.next[value]
		}
	}
	if temp.ending > 0 {
		return true
	}
	return false
}
