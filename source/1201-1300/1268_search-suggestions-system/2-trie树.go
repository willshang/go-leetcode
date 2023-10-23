package main

import "fmt"

func main() {
	fmt.Println(suggestedProducts([]string{"abc", "abc", "abc", "abd"}, "ab"))
}

func suggestedProducts(products []string, searchWord string) [][]string {
	res := make([][]string, len(searchWord))
	t := &Trie{}
	for i := 0; i < len(products); i++ {
		t.Insert(products[i])
	}
	for i := 0; i < len(searchWord); i++ {
		res[i] = t.StartsWith(searchWord[:i+1])
	}
	return res
}

type Trie struct {
	next [26]*Trie // 下一级指针，如不限于小写字母，[26]=>[256]
	str  string
}

// 插入word
func (this *Trie) Insert(word string) {
	temp := this
	for _, v := range word {
		value := v - 'a'
		if temp.next[value] == nil {
			temp.next[value] = &Trie{
				next: [26]*Trie{},
			}
		}
		temp = temp.next[value]
	}
	temp.str = word
}

// 查找前缀
func (this *Trie) StartsWith(prefix string) []string {
	temp := this
	for _, v := range prefix {
		value := v - 'a'
		if temp = temp.next[value]; temp == nil {
			return nil
		}
	}
	res := make([]string, 0)
	temp.dfs(&res)
	return res
}

func (this *Trie) dfs(res *[]string) {
	if len(*res) == 3 {
		return
	}
	if this.str != "" {
		*res = append(*res, this.str)
	}
	for _, v := range this.next {
		if len(*res) == 3 {
			return
		}
		if v == nil {
			continue
		}
		v.dfs(res)
	}
}
