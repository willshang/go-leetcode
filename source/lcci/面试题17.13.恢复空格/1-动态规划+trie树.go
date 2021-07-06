package main

import (
	"fmt"
)

func main() {
	fmt.Println(respace([]string{"abc", "ab"}, "aaa"))
}

func respace(dictionary []string, sentence string) int {
	n := len(sentence)
	root := &Trie{
		next: [26]*Trie{},
	}
	for i := 0; i < len(dictionary); i++ {
		root.Insert(reverse(dictionary[i])) // 反序插入
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + 1 // 上一个长度+1
		cur := root
		for j := i; j >= 1; j-- {
			value := int(sentence[j-1] - 'a')
			if cur.next[value] == nil {
				break
			} else if cur.next[value].ending > 0 { // 找到，更新
				dp[i] = min(dp[i], dp[j-1])
			}
			if dp[i] == 0 {
				break
			}
			cur = cur.next[value]
		}
	}
	return dp[n]
}

func reverse(s string) string {
	arr := []byte(s)
	for i := 0; i < len(s)/2; i++ {
		arr[i], arr[len(s)-1-i] = arr[len(s)-1-i], arr[i]
	}
	return string(arr)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
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
