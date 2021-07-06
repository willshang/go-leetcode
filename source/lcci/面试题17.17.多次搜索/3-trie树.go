package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiSearch("mississippi", []string{"is", "ppi", "hi", "sis", "i", "ssippi"}))
	//fmt.Println(multiSearch("abc", []string{""}))
}

func multiSearch(big string, smalls []string) [][]int {
	n := len(smalls)
	res := make([][]int, n)
	root := &Trie{
		next: [26]*Trie{},
	}
	for i := 0; i < n; i++ {
		root.Insert(smalls[i], i+1)
	}
	for i := 0; i < len(big); i++ {
		temp := root.Search(big[i:])
		for j := 0; j < len(temp); j++ {
			res[temp[j]] = append(res[temp[j]], i)
		}
	}
	return res
}

type Trie struct {
	next   [26]*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
	ending int       // 下标，从1开始
}

// 插入word
func (this *Trie) Insert(word string, index int) {
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
	temp.ending = index
}

// 查找
func (this *Trie) Search(word string) []int {
	arr := make([]int, 0) // 存放匹配到的下标列表
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return arr
		}
		if temp.ending > 0 {
			arr = append(arr, temp.ending-1)
		}
	}
	return arr
}
