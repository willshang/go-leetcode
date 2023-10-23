package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findAllConcatenatedWordsInADict([]string{"cat", "cats", "catsdogcats", "dog",
		"dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}))
}

// leetcode472_连接词
func findAllConcatenatedWordsInADict(words []string) []string {
	res := make([]string, 0)
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	root := Trie{}
	for i := 0; i < len(words); i++ {
		if words[i] == "" {
			continue
		}
		if root.Dfs(words[i]) == true {
			res = append(res, words[i])
		} else {
			root.Insert(words[i])
		}
	}
	return res
}

type Trie struct {
	next   [26]*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
	ending int       // 次数（可以改为bool）
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

func (this *Trie) Dfs(word string) bool {
	if word == "" {
		return true
	}
	temp := this
	for i := 0; i < len(word); i++ {
		value := word[i] - 'a'
		temp = temp.next[value]
		if temp == nil {
			return false
		}
		if temp.ending > 0 && this.Dfs(word[i+1:]) == true {
			return true
		}
	}
	return false
}
