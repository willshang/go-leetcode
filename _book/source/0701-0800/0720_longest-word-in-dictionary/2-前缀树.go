package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	//fmt.Println(longestWord(str))
	str_1 := []string{"worlds", "w", "wo", "wor", "worl", "world", "a", "banana", "app", "appl", "ap", "apply", "apple"}
	fmt.Println(longestWord(str_1))
}

type Trie struct {
	children [26]*Trie
	index    int
}

func Constructor(str string) Trie {
	return Trie{}
}

func (t *Trie) insert(str string, index int) {
	cur := t
	for i := 0; i < len(str); i++ {
		j := str[i] - 'a'
		if cur.children[j] == nil {
			cur.children[j] = &Trie{}
		}
		cur = cur.children[j]
	}
	cur.index = index
}

func (t *Trie) bfs(words []string) string {
	res := ""
	stack := make([]*Trie, 0)
	stack = append(stack, t)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.index > 0 || node == t {
			if node != t {
				word := words[node.index-1]
				if len(word) > len(res) || (len(word) == len(res) && word < res) {
					res = word
				}
			}
			for i := 0; i < len(node.children); i++ {
				if node.children[i] != nil {
					stack = append(stack, node.children[i])
				}
			}
		}
	}
	return res
}

func longestWord(words []string) string {
	t := Trie{}
	for i := 0; i < len(words); i++ {
		t.insert(words[i], i+1)
	}
	return t.bfs(words)
}
